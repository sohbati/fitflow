import {Component, ElementRef, OnInit, QueryList, ViewChild} from '@angular/core';

import {PersonService} from '../../../services/person.service';
import {HelperService} from '../../../services/helper.service';
import {Person} from '../../../datamodel/Person';
import {NgbActiveModal} from '@ng-bootstrap/ng-bootstrap';
import {PersonView} from '../../../datamodel/PersonView';
@Component({
  selector: 'ngx-add-person-modal',
  templateUrl: './add-person-modal.component.html',
  styleUrls: ['./add-person-modal.component.css'],
})

export class AddPersonModalComponent implements OnInit {

  public personId: number = 0;
  person: PersonView = new PersonView();

  constructor(private personService: PersonService,
              private helperService: HelperService,
              private ngbActiveModal: NgbActiveModal) {

  }

  ngOnInit() {
    if (this.personId === 0) {
      this.person.firstName = '';
      this.person.lastName = '';
      this.person.originalImage = './assets/images/dummy-person.png';
    }else {
      this.personService.getPersonById(this.personId).subscribe((person: PersonView) => {
        this.person = person;
        this.manageImage(this.person);
        if (this.person.birthDate === '1348/10/11') {
          this.person.birthDate = '';
        }
      }, error => {
        this.helperService.showError('خطا در نمایش اطلاعات فرد' + ' : ' + error);
      })
    }
  }

  @ViewChild('fileInput') inputEl: ElementRef;
  @ViewChild('imagePic') imagePic: ElementRef;

  showUserPicture(): void {
    const imagePic: HTMLInputElement = this.imagePic.nativeElement;
    const inputEl: HTMLInputElement = this.inputEl.nativeElement;
    if (inputEl.files && inputEl.files[0]) {
      const imgSize = inputEl.files[0].size;
      const MAX_IMAGE_SIZE = 3;
      if (imgSize > 1024 * 1024 * 3) {
        this.helperService.showError('عکس انتخاب شده نباید بیشتر از ۳ مگابایت باشد');
        return;
      }
      const reader = new FileReader();

      reader.onload = function (e: any) {
        const target: any = e.target;
        imagePic.src = target.result;
        imagePic.width = '150px';
        imagePic.height = '150px';
        imagePic.style.display = '';
      };

      reader.readAsDataURL(inputEl.files[0]);
    }
  }

  save() {
    const model = this.person;

    const inputEl: HTMLInputElement = this.inputEl.nativeElement;
    const formData: FormData = new FormData();
    const file = inputEl.files.item(0);

    if (file !== null) {
      formData.append('pic', file, 'filename');
      // formData.set('enctype', 'multipart/form-data')
    }
    // formData.append('person', JSON.stringify(model));
    formData.append('person', new Blob([JSON.stringify(model)],
      {
        type: 'application/json',
      }));

      this.personService.addOrUpdatePerson(formData, file !== null).subscribe((person: PersonView) => {
          this.person = person;
          this.helperService.showSuccess('اطلاعات ذخیره گردید');
          this.ngbActiveModal.close(this.person);
        }, err => {
          this.helperService.showError(err);
        },
      );
  }

  manageImage(person: PersonView) {
    const img = person.originalImage;
    if (img === undefined || img === null || img.length === 0) {
      person.originalImage = 'assets/images/dummy-person.png';
    }else {
      person.originalImage = this.helperService.BASE_64_IMAGE_PREFIX + person.originalImage;
    }
  }

  close() {
    this.ngbActiveModal.close();
  }
}
