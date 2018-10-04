import { Component, OnInit } from '@angular/core';
import { LocalDataSource } from 'ng2-smart-table';
import {ExerciseService} from '../../../../../services/exercise.service';
import {HelperService} from '../../../../../services/helper.service';
import {Exercise} from '../../../../../datamodel/Exercise';
import {NgbActiveModal} from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'ngx-add-exercise-modal',
  templateUrl: './add-exercise-modal.component.html',
  styleUrls: ['./add-exercise-modal.component.css'],
})

export class AddExerciseModalComponent implements OnInit {

  selectedExercise: Exercise = null;
  settings = {
    actions: {
      add: false,
      edit: false,
      delete: false,
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
              private ngbActiveModal: NgbActiveModal,
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

  selectExercise() {
    this.ngbActiveModal.close(this.selectedExercise);
  }

  closeClick() {
    this.ngbActiveModal.close();
  }

  lastClickTime: number = 0;

  onUserRowSelect(event: any) {
    if (this.lastClickTime === 0) {
      this.lastClickTime = new Date().getTime();
    } else {
      const change = (new Date().getTime()) - this.lastClickTime;
      if (change < 400) {
        this.onDoubleClick(event.data);
      }
      this.selectedExercise = event.data;
      this.lastClickTime = 0;
    }
  }

  onDoubleClick(row: Exercise) {
    this.selectedExercise = row;
    this.ngbActiveModal.close(row);
  }
}
