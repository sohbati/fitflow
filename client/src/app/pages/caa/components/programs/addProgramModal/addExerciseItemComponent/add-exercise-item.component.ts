import { Component, OnInit } from '@angular/core';
import { LocalDataSource } from 'ng2-smart-table';
import {CompleterService, CompleterData, CompleterItem} from 'ng2-completer';
import {Exercise} from '../../../../datamodel/Exercise';
import {ExerciseService} from '../../../../services/exercise.service';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import {HelperService} from '../../../../services/helper.service';
import {NotificationsService, NotificationType} from 'angular2-notifications';

interface SearchData {
  item: string;
  value: any;
}


@Component({
  selector: 'ngx-add-exercise-item-modal',
  templateUrl: './add-exercise-item.component.html',
  styleUrls: ['./add-exercise-item.component.css'],
})

export class AddExerciseItemModalComponent implements OnInit {
  source: LocalDataSource = new LocalDataSource();
  public dataService: CompleterData;
  public inputExercise: string;
  public inputExerciseValue: number;

  public inputSet: number;
  public inputRepeat: number;
  public inputRepeatType: string;
  public inputRowNumber: string;
  public searchData: SearchData[] = [];

  public exerciseRepeatTypeDesc: string[] = this.exerciseService.exerciseRepeatTypeDesc;

  public selectedExcercisesList: SelectedIExerciseItems[] = [];

  constructor(private exerciseService: ExerciseService,
              private ngbActiveModal: NgbActiveModal ,
              private helperService: HelperService,
              private completerService: CompleterService,
              private notificationService: NotificationsService) {
    this.dataService = completerService.local(this.searchData, 'value', 'item');
  }
  ngOnInit() {
    this.exerciseService.getExerciseShortList().subscribe((data: Exercise[]) => {
      data.forEach((e, index, array) => {
        this.searchData.push({item : e.name, value : e.id});
      });
    });

    if (this.selectedExcercisesList != null && this.selectedExcercisesList.length > 0) {
      this.selectedExcercisesList.forEach((value, index, array) => {
        value.subListItems.forEach((value2, index2, array2) => {
          value2.repeatTypeDesc = this.exerciseService.getExerciseRepeatTypeDesc(value2.repeatType);
        });
      });
    }
  }

  validateAdd(): boolean {
    let r = true;
    if ((this.inputExercise == null || this.inputExercise === '') || (this.inputSet == null) ||
      (this.inputRepeat == null)) {
      r = false;
    }
    let num: number = parseInt(this.helperService.convertToLatinNumbers(this.inputRowNumber));
    if (isNaN(num)) {
      this.helperService.showError('لطفا شماره ردیف را عددی وارد کنید');
      return false;
    }
    num = parseInt(this.helperService.convertToLatinNumbers(this.inputSet + ''));
    if (isNaN(num)) {
      this.helperService.showError('لطفا  تعداد ست را عددی وارد کنید');
      return false;
    }
    num = parseInt(this.helperService.convertToLatinNumbers(this.inputRepeat + ''));
    if (isNaN(num)) {
      this.helperService.showError('لطفا تعداد تکرار را عددی وارد کنید');
      return false;
    }
    return r;
  }
  addClick() {
    if (!this.validateAdd()) {
      return;
    }
    let item: SelectedIExerciseItems ;
    if (this.selectedExcercisesList == null) {
      this.selectedExcercisesList = [];
    }
    const rowNum: number = (
      this.inputRowNumber == null  ? this.selectedExcercisesList.length :
        parseInt(this.helperService.convertToLatinNumbers(this.inputRowNumber)));
    let isNewRow = true;
    if (this.selectedExcercisesList.length === 0 || rowNum > this.selectedExcercisesList.length) {
      item = {
        id : this.selectedExcercisesList.length + 1,
        subListItems : [],
      }
    }else {
      for (let i = 0; i < this.selectedExcercisesList.length; i++) {
        if (this.selectedExcercisesList[i].id === rowNum) {
          item = this.selectedExcercisesList[i];
          isNewRow = false;
          break;
        }
      }
    }
    const subItem: ExerciseItemsSubList = {
      exerciseItemDesc: this.inputExercise,
      exerciseItemId: this.inputExerciseValue,
      repeat: parseInt(this.helperService.convertToLatinNumbers(this.inputRepeat + '') + ''),
      repeatType: this.inputRepeatType,
      repeatTypeDesc: this.exerciseService.getExerciseRepeatTypeDesc(this.inputRepeatType),
      set: parseInt(this.helperService.convertToLatinNumbers(this.inputSet + '') + ''),
    }

    if (item.subListItems == null) {
      item.subListItems = [];
    }
    item.subListItems.push(subItem);
    if (isNewRow) {
      this.selectedExcercisesList.push(item);
    }
  }

  removeItem (event, index: number) {
    this.selectedExcercisesList.splice(index - 1, 1);
    for (let i = 0; i < this.selectedExcercisesList.length; i++) {
      this.selectedExcercisesList[i].id = i + 1;
    }
  }
  completerSelected(event: CompleterItem) {
    this.inputExerciseValue = event.originalObject.value;
  }

  confirmClick() {
    this.ngbActiveModal.close(this.selectedExcercisesList);
  }

}
