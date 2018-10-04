import { Component, OnInit } from '@angular/core';
import { LocalDataSource } from 'ng2-smart-table';
import {Exercise} from '../../../../datamodel/Exercise';
import {ExerciseService} from '../../../../services/exercise.service';
import {NgbActiveModal, NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {HelperService} from '../../../../services/helper.service';
import {NotificationsService, NotificationType} from 'angular2-notifications';
import DataSource from 'devextreme/data/data_source';
import ArrayStore from 'devextreme/data/array_store';
import {AddExerciseModalComponent} from './addExerciseModalComponent/add-exercise-modal.component';

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
  /**
   * used in dxAutoCompleter
   * this code initialized again in init method  **/
  dataSource =  new DataSource({
    store: new ArrayStore({
      key: 'item',
      data: [],
    }),
  });

  source: LocalDataSource = new LocalDataSource();
  public inputExercise: string;
  public inputExerciseId: number = 0;


  public inputSet: number;
  public inputRepeat: number;
  public inputRepeatType: string;
  public inputRowNumber: string;
  public searchData: SearchData[] = [];

  public exerciseRepeatTypeDesc: string[] = this.exerciseService.exerciseRepeatTypeDesc;

  public selectedExcercisesList: SelectedIExerciseItems[] = [];

  constructor(private modalService: NgbModal,
              private exerciseService: ExerciseService,
              private ngbActiveModal: NgbActiveModal ,
              private helperService: HelperService,
              private notificationService: NotificationsService) {
  }
  ngOnInit() {
    const programList = [];
    this.exerciseService.getExerciseShortList().subscribe((data: Exercise[]) => {
      data.forEach((e, index, array) => {
        this.searchData.push({item : e.name, value : e.id});
        programList.push(e.name);
      });

      this.dataSource =  new DataSource({
        store: new ArrayStore({
          key: 'item',
          data: programList,
        }),
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
    let num: number = this.helperService.toInt(this.helperService.convertToLatinNumbers(this.inputRowNumber));
    if (isNaN(num)) {
      this.helperService.showError('لطفا شماره ردیف را عددی وارد کنید');
      return false;
    }
    num = this.helperService.toInt(this.helperService.convertToLatinNumbers(this.inputSet + ''));
    if (isNaN(num)) {
      this.helperService.showError('لطفا  تعداد ست را عددی وارد کنید');
      return false;
    }
    num = this.helperService.toInt(this.helperService.convertToLatinNumbers(this.inputRepeat + ''));
    if (isNaN(num)) {
      this.helperService.showError('لطفا تعداد تکرار را عددی وارد کنید');
      return false;
    }

    const t = this.inputRepeatType;
    if (t === '' || t.length === 0) {
      this.helperService.showError('نوع تکرار نمی تواند خالی باشد');
      return false;
    }
    return r;
  }

  getExerciseId(name: string) {
    let exercixeId = 0;
    this.searchData.forEach((programItem: SearchData, index, array) => {
      if (programItem.item === name) {
        exercixeId = programItem.value;
      }
    });
    return exercixeId;
  }
  validateSuperItemSet(subListItems: ExerciseItemsSubList[] , inputSet: number) {
    const set = this.helperService.convertToLatinNumbers(inputSet + '');
    if (set == null || set.length === 0) {
      return true;
    }

    if (subListItems == null || subListItems.length === 0) {
      return true;
    }
    let found = false;
    for (let i = 0; i < subListItems.length; i++) {
      if (subListItems[i].set === this.helperService.toInt(set)) {
         found = true;
      }
    }
    if (!found) {
      this.helperService.showError('برای حرکت سوپر تعداد ست ها نمی تواند متفاوت باشد');
      return false;
    }
    return true;
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
        this.helperService.toInt(this.helperService.convertToLatinNumbers(this.inputRowNumber)));
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
          if (!this.validateSuperItemSet(item.subListItems, this.inputSet)) {
            return;
          }
          isNewRow = false;
          break;
        }
      }
    }
    if (this.inputExerciseId <= 0 && this.inputExercise.length > 0) {
      this.inputExerciseId = this.getExerciseId(this.inputExercise);
    }
    const exerciseId = this.inputExerciseId;
    if (exerciseId === 0) {
      this.helperService.showError('خطای پیاده ساز. شماره حرکت پیدا نشد!!!')
      return;
    }

    const subItem: ExerciseItemsSubList = {
      exerciseItemDesc: this.inputExercise,
      exerciseItemId: exerciseId,
      repeat: this.helperService.toInt(this.helperService.convertToLatinNumbers(this.inputRepeat + '') + ''),
      repeatType: this.inputRepeatType,
      repeatTypeDesc: this.exerciseService.getExerciseRepeatTypeDesc(this.inputRepeatType),
      set: this.helperService.toInt(this.helperService.convertToLatinNumbers(this.inputSet + '') + ''),
    }

    if (item.subListItems == null) {
      item.subListItems = [];
    }
    item.subListItems.push(subItem);
    if (isNewRow) {
      this.selectedExcercisesList.push(item);
    }

    this.inputExercise = '';
    this.inputSet = null;
    this.inputRepeat = null;
    this.inputRepeatType = '';
  }

  removeItem (event, index: number) {
    this.selectedExcercisesList.splice(index - 1, 1);

    this.reindexSelectedExcercisesList();
  }

  confirmClick() {
    this.ngbActiveModal.close(this.selectedExcercisesList);
  }

  exerciseModalOpenclick() {
    const activeModal = this.modalService.open(AddExerciseModalComponent, { backdrop: 'static', size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.modalHeader = 'Large Modal';
    activeModal.result.then((result: Exercise) => {
      this.inputExercise = result.name;
      this.inputExerciseId = result.id;
    });

  }

  closeClick() {
    this.ngbActiveModal.close(null);
  }

  moveUpItem(event: any, index: number) {
    if (index === 1) {
      return;
    }

    const tmp = this.selectedExcercisesList[index - 1];
    this.selectedExcercisesList[index - 1] = this.selectedExcercisesList[index - 2];
    this.selectedExcercisesList[index - 2] = tmp;

    this.reindexSelectedExcercisesList();
  }

  moveDownItem(event: any, index: number) {
    if (index === this.selectedExcercisesList.length) {
      return;
    }

    const tmp = this.selectedExcercisesList[index - 1];
    this.selectedExcercisesList[index - 1] = this.selectedExcercisesList[index];
    this.selectedExcercisesList[index] = tmp;

    this.reindexSelectedExcercisesList();
  }

  reindexSelectedExcercisesList() {
    for (let i = 0; i < this.selectedExcercisesList.length; i++) {
      this.selectedExcercisesList[i].id = i + 1;
    }
  }
}
