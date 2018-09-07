import { Component, OnInit } from '@angular/core';
import {PersonService} from '../../services/person.service';
import {HelperService} from '../../services/helper.service';
import {Person} from '../../datamodel/Person';
import { LocalDataSource } from 'ng2-smart-table';
@Component({
  selector: 'ngx-person-list',
  templateUrl: './persons.component.html',
  styleUrls: ['./persons.component.css'],
})

export class PersonsComponent implements OnInit {


  settings = {
    add: {
      addButtonContent: '<i class="nb-plus"></i>',
      createButtonContent: '<i class="nb-checkmark"></i>',
      cancelButtonContent: '<i class="nb-close"></i>',
      confirmCreate: true,
    },
    edit: {
      editButtonContent: '<i class="nb-edit"></i>',
      saveButtonContent: '<i class="nb-checkmark"></i>',
      cancelButtonContent: '<i class="nb-close"></i>',
      confirmSave: true,
    },
    delete: {
      deleteButtonContent: '<i class="nb-trash"></i>',
      confirmDelete: true,
    },
    columns: {
      firstName: {
        title: 'نام ',
        type: 'string',
      },
      lastName: {
        title: 'نام خانوادگی',
        type: 'string',
      },
      mobileNumber: {
        title: 'شماره موبایل',
        type: 'string',
      },
    },
  };

  source: LocalDataSource = new LocalDataSource();


  constructor(private personService: PersonService,
              private helperService: HelperService) {

    const data =  [{
      id: 1,
      firstName: 'Mark',
      lastName: 'Otto',
      username: '@mdo',
      email: 'mdo@gmail.com',
      age: '28',
    }, {
      id: 2,
      firstName: 'Jacob',
      lastName: 'Thornton',
      username: '@fat',
      email: 'fat@yandex.ru',
      age: '45',
    }, {
      id: 3,
      firstName: 'Larry',
      lastName: 'Bird',
      username: '@twitter',
      email: 'twitter@outlook.com',
      age: '18',
    }];

    //this.source.load(data);

  }

  private personList: Person[] = [];
  ngOnInit() {
      this.initPersonList();
  }

  initPersonList() {
    this.personService.getPersonList().subscribe((data: Person[]) => {
      console.log(data);
      this.personList = data;
      this.source.load(this.personList);
    });
  }

  onDeleteConfirm(event): void {
    if (window.confirm('Are you sure you want to delete?')) {
      this.doDelete(event);
    } else {
      event.confirm.reject();
    }
  }

  doDelete(event): void {
    this.personService.deletePerson(event.data).subscribe((data: any) => {
      event.confirm.resolve();
      this.helperService.showSuccess('اطلاعات حذف گردید');
    }, (error) => {
      this.helperService.showError('اگر برای  فرد برنامه ای تنظیم شده است نمی توان آن شخص را حذف نمود')
      event.confirm.reject()
    });
  }

  onEditConfirm(event): void {
    console.log(event);
    this.personService.editPerson(event.newData).subscribe((data: any) => {
      event.confirm.resolve();
      this.helperService.showSuccess('اطلاعات ذخیره گردید');
    }), error => {
      console.log('Error' + error);
      this.helperService.showError('خطا در ذخیره اطلاعات در بانک اطلاعاتی')
      event.confirm.reject()
    };
  }

  onCreateConfirm(event): void {
    console.log(event);
    this.personService.addPerson(event.newData).subscribe((data: any) => {
      event.confirm.resolve();
      this.helperService.showSuccess('اطلاعات ذخیره گردید');
    }), error => {
      console.log('Error' + error);
      this.helperService.showError('خطا در ذخیره اطلاعات در بانک اطلاعاتی')
      event.confirm.reject()
    };
  }
}
