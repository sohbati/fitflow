import { Component, OnInit } from '@angular/core';
import {ExerciseService} from '../../services/exercise.service';
import {HelperService} from '../../services/helper.service';
import {Exercise} from '../../datamodel/Exercise';
import { LocalDataSource } from 'ng2-smart-table';

@Component({
  selector: 'ngx-exercise-list',
  templateUrl: './exercises.component.html',
  styleUrls: ['./exercises.component.css'],
})

export class ExercisesComponent implements OnInit {


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
      name: {
        title: 'نام ',
        type: 'string',
      },
      code: {
        title: 'کد حرکت',
        type: 'string',
      },
      latinName: {
        title: 'نام لاتین',
        type: 'string',
      },
      involvedMuscel: {
        title: 'عضله درگیر',
        type: 'string',
      },
    },
  };

  source: LocalDataSource = new LocalDataSource();


  constructor(private exerciseService: ExerciseService,
              private helperService: HelperService) {
  }

  private exerciseList: Exercise[] = [];
  ngOnInit() {
      this.initExerciseList();
  }

  initExerciseList() {
    this.exerciseService.getExerciseList().subscribe((data: Exercise[]) => {
      this.exerciseList = data;
      this.source.load(this.exerciseList);
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
    this.exerciseService.deleteExercise(event.data).subscribe((data: any) => {
      event.confirm.resolve();
      this.helperService.showSuccess('اطلاعات حذف گردید');
    }, error => {
      this.helperService.showError('error:' )
      event.confirm.reject()
    });
  }

  onEditConfirm(event): void {
    this.exerciseService.editExercise(event.newData).subscribe((data: any) => {
      event.confirm.resolve();
      this.helperService.showSuccess('اطلاعات ذخیره گردید');
    }, error => {
      this.helperService.showError2('خطا در ذخیره اطلاعات در بانک اطلاعاتی', error)
      event.confirm.reject()
    });
  }

  onCreateConfirm(event): void {
    this.exerciseService.addExercise(event.newData).subscribe((data: Exercise) => {
      event.confirm.resolve(data);

      this.helperService.showSuccess('اطلاعات ذخیره گردید');
    }, error => {
      this.helperService.showError2('خطا در ذخیره اطلاعات در بانک اطلاعاتی', error)
      event.confirm.reject()
    });
  }
}
