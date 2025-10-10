import { Component, OnInit } from '@angular/core';
import {ExerciseService} from '../../services/exercise.service';
import {HelperService} from '../../services/helper.service';
import {Exercise} from '../../datamodel/Exercise';
import { LocalDataSource } from 'ng2-smart-table';
import {LangChangeEvent, TranslateService} from '@ngx-translate/core';

@Component({
  selector: 'ngx-exercise-list',
  templateUrl: './exercises.component.html',
  styleUrls: ['./exercises.component.css'],
})

export class ExercisesComponent implements OnInit {
  settings = {};
  source: LocalDataSource = new LocalDataSource();

  constructor(private exerciseService: ExerciseService,
              private helperService: HelperService,
              private translate: TranslateService) {
  }

  private exerciseList: Exercise[] = [];
  ngOnInit() {
    this.translate.use('fa').subscribe((event: LangChangeEvent) => {
      this.settings = {
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
            title: this.translate.instant('EXERCISE.NAME'),
              type: 'string',
          },
          code: {
            title: this.translate.instant('EXERCISE.CODE'),
              type: 'string',
          },
          latinName: {
            title: this.translate.instant('EXERCISE.LATIN_NAME'),
              type: 'string',
          },
          involvedMuscel: {
            title: this.translate.instant('EXERCISE.INVOLVED_MUSCLE'),
              type: 'string',
          },
        },
      };
    });
    this.initExerciseList();
  }

  initExerciseList() {
    this.exerciseService.getExerciseList().subscribe((data: Exercise[]) => {
      this.exerciseList = data;
      this.source.load(this.exerciseList);
    }, err => {
      this.helperService.showError(err);
    });
  }

  onDeleteConfirm(event): void {
    const msg = this.translate.instant('PUBLIC.DELETE_CONFIRM_MSG');
    if (window.confirm( msg)) {
      this.doDelete(event);
    } else {
      event.confirm.reject();
    }
  }

  doDelete(event): void {
    this.exerciseService.deleteExercise(event.data).subscribe((data: any) => {
      event.confirm.resolve();
      this.helperService.showSuccess(this.translate.instant('PUBLIC.DELETE_DONE'));
    }, error => {
      this.helperService.showError(error);
      event.confirm.reject()
    });
  }

  onEditConfirm(event): void {
    this.exerciseService.editExercise(event.newData).subscribe((data: any) => {
      event.confirm.resolve();
      this.helperService.showSuccess(this.translate.instant('PUBLIC.SAVE_DONE'));
    }, error => {
      this.helperService.showError2(this.translate.instant('PUBLIC.ERROR_IN_SAVING'), error)
      event.confirm.reject()
    });
  }

  onCreateConfirm(event): void {
    this.exerciseService.addExercise(event.newData).subscribe((data: Exercise) => {
      event.confirm.resolve(data);

      this.helperService.showSuccess(this.translate.instant('PUBLIC.SAVE_DONE'));
    }, error => {
      this.helperService.showError2(this.translate.instant('PUBLIC.ERROR_IN_SAVING'), error)
      event.confirm.reject()
    });
  }
}
