import { Component, OnInit } from '@angular/core';
import {PersonService} from '../../services/person.service';
import {HelperService} from '../../services/helper.service';
import {Person} from '../../datamodel/Person';
import {AddPersonModalComponent} from './addpersonmodal/add-person-modal.component';
import {NgbActiveModal, NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {PersonView} from '../../datamodel/PersonView';
import {ProgramsComponent} from '../programs/programs.component';
@Component({
  selector: 'ngx-person-list',
  templateUrl: './persons.component.html',
  styleUrls: ['./persons.component.css'],
})

export class PersonsComponent implements OnInit {


  constructor(private personService: PersonService,
              private helperService: HelperService,
              private modalService: NgbModal) {
  }

  searchText: string = '';
  private personList: Person[] = [];

  ngOnInit() {
      this.initPersonList();
  }
  ///////////////////////////////////////////////////////////////////////////
  initPersonList() {
    this.personService.getPersonList().subscribe((data: Person[]) => {
      this.personList = data;
      this.personList.forEach((person, index, array) => {
        this.manageImage(person);
      });
    });
  }

  manageImage(person: PersonView) {
    const img = person.shrinkedImage;
    if (img === undefined || img === null || img.length === 0)
      person.shrinkedImage = 'assets/images/dummy-person.png';
    else
      person.shrinkedImage = this.helperService.BASE_64_IMAGE_PREFIX + person.shrinkedImage;
    // return person;
  }

  addPersonClick() {
    const activeModal = this.modalService.open(AddPersonModalComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.personId = 0;
    activeModal.componentInstance.modalHeader = 'Large Modal';

    activeModal.result.then((result: any) => {
      this.manageImage(result);
      this.personList.push(result);
    }).catch(e => {
      // this.helperService.showError('error on modal edit program : ' + e)
    });
  }

  editButtonClick(id: number) {
    const activeModal = this.modalService.open(AddPersonModalComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.personId = id;
    activeModal.componentInstance.modalHeader = 'Large Modal';

    activeModal.result.then((result: PersonView) => {
      for (let i = 0; i < this.personList.length; i++) {
          if (this.personList[i].id === id) {
            this.manageImage(result);
            this.personList[i] = result;
            break;
          }
      }
    }).catch(e => {
      // this.helperService.showError('error on modal edit program : ' + e)
    });
  }

  OpenProgramClick(person: PersonView) {
    const activeModal = this.modalService.open(ProgramsComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.person = person;
    activeModal.componentInstance.modalHeader = 'Large Modal';

    activeModal.result.then((result: PersonView) => {

    }).catch(e => {
      // this.helperService.showError('error on modal edit program : ' + e)
    });
  }
  deleteClick(id: number): void {
    if (window.confirm('آیا از حذف فرد مورد نظر اطمینان دارید؟')) {
      this.doDelete(id);
    }
  }

  doDelete(id: number): void {
    this.personService.deletePerson(id).subscribe((data: any) => {
      let index = -1;
      for (let i = 0; i < this.personList.length; i++) {
        if (this.personList[i].id === id) {
          index = i;
          break;
        }
      }
      if (index >= 0) {
        this.personList.splice(index, 1);
      }
      this.helperService.showSuccess('اطلاعات حذف گردید');
    }, (error) => {
      this.helperService.showError('اگر برای  فرد برنامه ای تنظیم شده است نمی توان آن شخص را حذف نمود')
    });
  }


  public searchClick() {
    let s = this.searchText;
    if (s ===  '' || s.length === 0 ) {
      s = 'EMPTY';
    }
    this.personService.findByNameFamilyPhone(s).subscribe((list: PersonView[]) => {
        this.personList = list;
        if (this.personList && this.personList.length > 0) {
          this.personList.forEach((person, index, array) => {
            this.manageImage(person);
          });
        }
    }, err => {
      this.helperService.showError2(' خطا در جستجو', err);
    })

  }
}
