import { Component, OnInit } from '@angular/core';
import {HelperService} from '../../../services/helper.service';
import { LocalDataSource } from 'ng2-smart-table';
import {AddExerciseItemModalComponent} from './addExerciseItemComponent/add-exercise-item.component';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {ProgramService} from '../../../services/program.service';
import {ProgramExerciseItem, ProgramView} from '../../../datamodel/ProgramView';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import {ExerciseService} from '../../../services/exercise.service';
import {DisplayProgramExerciseImageComponent} from './displayProgramExerciseImageComponent/display-program-exercise-image.component';
import {DisplayAllSizesForOnePersonComponent} from './displayAllSizesForOnePerson/display-all-sizes-for-one-person.component';
import {PersonView} from '../../../datamodel/PersonView';


interface ExerciseAsString {
  rowId: number;
  exercises: string;
}

@Component({
  selector: 'ngx-add-program-modal',
  templateUrl: './add-program-modal.component.html',
  styleUrls: ['./add-program-modal.component.css'],
})

/**
 *
 */
export class AddProgramModalComponent implements OnInit {


  /** ng2-smart-table source **/
  source1: LocalDataSource = new LocalDataSource();
  source2: LocalDataSource = new LocalDataSource();
  source3: LocalDataSource = new LocalDataSource();
  source4: LocalDataSource = new LocalDataSource();

  exercise1List: SelectedIExerciseItems[] = [];
  exercise2List: SelectedIExerciseItems[] = [];
  exercise3List: SelectedIExerciseItems[] = [];
  exercise4List: SelectedIExerciseItems[] = [];

  program: ProgramView;
  person: PersonView;

  /** ng2-smart-table setting **/
  settings = {
    mode: 'external',
    actions: {
      columnTitle: 'Actions',
      add: true,
      edit: false,
      delete: false,
    },
    add: {
      addButtonContent: '<i class="nb-plus"></i>',
      createButtonContent: '<i class="nb-checkmark"></i>',
      cancelButtonContent: '<i class="nb-close"></i>',
      // confirmCreate: true,
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
      rowId: {
        title: 'ردیف',
        type: 'string',
        filter: false,
        width: '6%',
        editable: false,
      },
      exercises: {
        title: 'تمرینات ',
        type: 'string',
        filter: true,
        width: '94%',
        editable: false,
      },
    },
  };

  constructor(private modalService: NgbModal,
              private ngbActiveModal: NgbActiveModal,
        private helperService: HelperService,
        private programService: ProgramService,
        private exerciseService: ExerciseService) {
  }
  ngOnInit() {
    if (this.program == null) {
      this.program = new ProgramView();
      this.program.programName = '';
      this.program.shamsiProgramDate = '';
      this.program.description = '';
      this.program.personAge = 0;
      this.program.personTall = 0;
      this.program.personWeight = 0;
      this.program.personChest = 0;
      this.program.personWaist = 0;
      this.program.personAbdomen = 0;
      this.program.personArm = 0;
      this.program.personForeArm = 0;
      this.program.personThigh = 0;
      this.program.personShin = 0;
      this.program.personButt = 0;
      this.program.personFatPercentage = 0;
      this.program.personFatWeight = 0;
      this.program.personMuscleWeight = 0;
      this.program.personScore = 0;

      this.exercise1List = [];
      this.exercise2List = [];
      this.exercise3List = [];
      this.exercise4List = [];
    }else {
      this.intiExerciseItems(this.program.programExercise1Items, this.exercise1List, 1);
      this.intiExerciseItems(this.program.programExercise2Items, this.exercise2List, 2);
      this.intiExerciseItems(this.program.programExercise3Items, this.exercise3List, 3);
      this.intiExerciseItems(this.program.programExercise4Items, this.exercise4List, 4);
    }

    this.program.personName = this.person.firstName + ' ' + this.person.lastName;
    this.program.person = new PersonView();
    this.program.person.id = this.person.id;
    this.program.person.firstName = this.person.firstName;
    this.program.person.lastName = this.person.lastName;

  }

  intiExerciseItems(programExerciseItems: ProgramExerciseItem[],
                    exerciseList: SelectedIExerciseItems[], exerciseIndex: number) {
    if (programExerciseItems == null || programExerciseItems.length <= 0) {
      exerciseList = [];
    }else {
      programExerciseItems.forEach((item: ProgramExerciseItem, index, array) => {
        this.addItemToList(exerciseList, item.id, item)
      });
      exerciseList.forEach((value, index, array) => {
        value.id  = index + 1;
      })
    }
    this.prepareGridRows(exerciseList, exerciseIndex);
  }

  addItemToList(exerciseList: SelectedIExerciseItems[], id: number, item: ProgramExerciseItem) {
    if (exerciseList == null || exerciseList.length === 0) {
      this.addNewItem(exerciseList, item);
    }else {
      let isnew: boolean = true;
      for (let i = 0; i < exerciseList.length; i++) {
        if (exerciseList[i].id === id) {
          isnew = false;
          const subItem = this.getSubListItem(item);
          exerciseList[i].subListItems.push(subItem);
        }
      }
      if (isnew) {
        this.addNewItem(exerciseList, item);
      }
    }
  }
  addNewItem(exerciseList: SelectedIExerciseItems[], item: ProgramExerciseItem) {
    if (exerciseList == null) {
      exerciseList = [];
    }

    const subListItem: ExerciseItemsSubList = this.getSubListItem(item);
    const e: SelectedIExerciseItems = {
      id : item.id,
    }
    e.subListItems = [];
    e.subListItems.push(subListItem);
    exerciseList.push(e);
  }

  getSubListItem(item: ProgramExerciseItem): ExerciseItemsSubList {
    const subListItem: ExerciseItemsSubList = {
      exerciseItemDesc: item.exerciseName,
      exerciseItemId: item.exerciseId,
      set: item.exerciseSet,
      repeat: item.exerciseRepeat,
      repeatType: item.exerciseRepeatType,
      repeatTypeDesc: '',
    }
    return subListItem;
  }
////////////////////////////////////////////////////////////////////////////
  addExerciseClick(event, exerciseIndex: number): void {
    const activeModal = this.modalService.open(AddExerciseItemModalComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.selectedExcercisesList = this.getCurrentExercise(exerciseIndex);
    activeModal.componentInstance.modalHeader = 'Large Modal';

    activeModal.result.then((exerciseList: SelectedIExerciseItems[]) => {
      switch (exerciseIndex) {
        case 1 : this.exercise1List = exerciseList; break;
        case 2 : this.exercise2List = exerciseList; break;
        case 3 : this.exercise3List = exerciseList; break;
        case 4 : this.exercise4List = exerciseList; break;
      }
      this.prepareGridRows(exerciseList, exerciseIndex);
    }).catch(e => {
      // this.helperService.showError('error on modal add program : ' + e)
    });
  }

  getCurrentExercise(index: number) {
    switch (index) {
      case 1 : return this.exercise1List;
      case 2 : return this.exercise2List;
      case 3 : return this.exercise3List;
      case 4 : return this.exercise4List;
    }
  }

  editExerciseClick(event, exerciseIndex: number) {
    const activeModal = this.modalService.open(AddExerciseItemModalComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.selectedExcercisesList = this.getCurrentExercise(exerciseIndex);
    activeModal.componentInstance.modalHeader = 'Large Modal';

    activeModal.result.then((exerciseList: SelectedIExerciseItems[]) => {
      switch (exerciseIndex) {
        case 1 : this.exercise1List = exerciseList; break;
        case 2 : this.exercise2List = exerciseList; break;
        case 3 : this.exercise3List = exerciseList; break;
        case 4 : this.exercise4List = exerciseList; break;
      }

      this.prepareGridRows(exerciseList, exerciseIndex);
    }).catch(e => {
     // this.helperService.showError('error on modal edit program : ' + e)
    });
  }

  prepareGridRows(exerciseList: SelectedIExerciseItems[], exerciseIndex: number) {
    const exerciseAsStringList: ExerciseAsString[] = [];
    if (exerciseList != null && exerciseList.length > 0) {
      let str = '';
      for (let i = 0; i < exerciseList.length; i++) {
        const subArr = exerciseList[i].subListItems;
        str = '';
        subArr.forEach((item, index, array) => {
          const repeatTypedesc = this.exerciseService.getExerciseRepeatTypeDesc(item.repeatType);
          str += ' (' + item.exerciseItemDesc + ' ' + item.set + ' *' + item.repeat + ' ' + repeatTypedesc + ') '
        })
        exerciseAsStringList.push({rowId: exerciseList[i].id,  exercises: str});
      }
    }
    switch (exerciseIndex) {
      case 1 : this.source1.load(exerciseAsStringList); break;
      case 2 : this.source2.load(exerciseAsStringList); break;
      case 3 : this.source3.load(exerciseAsStringList); break;
      case 4 : this.source4.load(exerciseAsStringList); break;
    }
  }

  prepareExerciseList(exerciseList: SelectedIExerciseItems[]) {
    if (exerciseList == null || exerciseList.length === 0) {
      return [];
    }
    const programExerciseItems: ProgramExerciseItem[] = [];
    let id: number = 0;
    for (let i = 0; i < exerciseList.length; i++) {
      const item = exerciseList[i];
    // exerciseList.forEach((item, index, array) => {
      id++;
      for (let j = 0; j < item.subListItems.length; j++) {
        const value = item.subListItems[j];
      // item.subListItems.forEach((value, index2, array2) => {
        const ex: ProgramExerciseItem = new ProgramExerciseItem();
        ex.id = id;
        ex.exerciseId = value.exerciseItemId;
        ex.exerciseName = value.exerciseItemDesc;
        ex.exerciseSet = value.set;
        ex.exerciseRepeat = value.repeat;
        ex.exerciseRepeatType = value.repeatType;

        programExerciseItems.push(ex);
      }
    }
    return programExerciseItems;
  }
  validate(): boolean {
    if (this.program.programName == null || this.program.programName === '') {
      this.helperService.showError('لطفا نام برنامه را وارد نمایید')
      return false;
    }
    if (!this.validateNumberField(this.program.personAge, 'سن')) { return; }
    if (!this.validateNumberField(this.program.personTall, 'قد')) { return; }
    if (!this.validateNumberField(this.program.personWeight, 'وزن')) { return; }
    if (!this.validateNumberField(this.program.personChest, 'سینه')) { return; }
    if (!this.validateNumberField(this.program.personWaist, 'کمر')) { return; }
    if (!this.validateNumberField(this.program.personAbdomen, 'شکم')) { return; }
    if (!this.validateNumberField(this.program.personArm, 'بازو')) { return; }
    if (!this.validateNumberField(this.program.personForeArm, 'ساعد')) { return; }
    if (!this.validateNumberField(this.program.personThigh, 'ران')) { return; }
    if (!this.validateNumberField(this.program.personShin, 'ساق')) { return; }
    if (!this.validateNumberField(this.program.personButt, 'باسن')) { return; }
    if (!this.validateNumberField(this.program.personFatPercentage, 'درصد چربی')) { return; }
    if (!this.validateNumberField(this.program.personFatWeight, 'وزن چربی')) { return; }
    if (!this.validateNumberField(this.program.personMuscleWeight, 'وزن عضله')) { return; }
    if (!this.validateNumberField(this.program.personScore, 'امتیاز')) { return; }
    // if (!this.validateNumberField(this.program.personAge, 'سن')) { return; }

    return true;
  }
  validateNumberField(f: any, fieldLabel: string) {
    let num: number = parseInt(this.helperService.convertToLatinNumbers(f + ''));
    if (isNaN(num)) {
      this.helperService.showError('لطفا ' + fieldLabel + ' را عددی وارد کنید');
      return false;
    }
    return true;
  }
  confirmClick() {
    if (!this.validate()) {
      return;
    }
    this.program.personAge = parseInt(this.helperService.convertToLatinNumbers(this.program.personAge + ''));
    this.program.personTall = parseInt(this.helperService.convertToLatinNumbers(this.program.personTall + ''));
    this.program.personWeight = parseInt(this.helperService.convertToLatinNumbers(this.program.personWeight + ''));
    this.program.personChest = parseInt(this.helperService.convertToLatinNumbers(this.program.personChest + ''));
    this.program.personWaist = parseInt(this.helperService.convertToLatinNumbers(this.program.personWaist + ''));
    this.program.personAbdomen = parseInt(this.helperService.convertToLatinNumbers(this.program.personAbdomen + ''));
    this.program.personArm = parseInt(this.helperService.convertToLatinNumbers(this.program.personArm + ''));
    this.program.personForeArm = parseInt(this.helperService.convertToLatinNumbers(this.program.personForeArm + ''));
    this.program.personThigh = parseInt(this.helperService.convertToLatinNumbers(this.program.personThigh + ''));
    this.program.personShin = parseInt(this.helperService.convertToLatinNumbers(this.program.personShin + ''));
    this.program.personButt = parseInt(this.helperService.convertToLatinNumbers(this.program.personButt + ''));
    this.program.personFatPercentage = parseInt(this.helperService.convertToLatinNumbers(this.program.personFatPercentage + ''));
    this.program.personFatWeight = parseInt(this.helperService.convertToLatinNumbers(this.program.personFatWeight + ''));
    this.program.personMuscleWeight = parseInt(this.helperService.convertToLatinNumbers(this.program.personMuscleWeight + ''));
    this.program.personScore = parseInt(this.helperService.convertToLatinNumbers(this.program.personScore + ''));

    this.program.programExercise1Items = this.prepareExerciseList(this.exercise1List);
    this.program.programExercise2Items = this.prepareExerciseList(this.exercise2List);
    this.program.programExercise3Items = this.prepareExerciseList(this.exercise3List);
    this.program.programExercise4Items = this.prepareExerciseList(this.exercise4List);
    this.programService.addProgram(this.program).subscribe((result: ProgramView) => {
      this.program.id = result.id;
      this.helperService.showSuccess('اطلاعات با موفقیت ثبت گردید')
      this.ngbActiveModal.close(this.program);
    },
    error2 => {
       this.helperService.showError(' : خطا در ذخیره اطلاعات ' + error2);
    });
  }

  exportProgramClick() {
      const activeModal = this.modalService.open(DisplayProgramExerciseImageComponent,
        {windowClass : 'modalWindowClass', container: 'nb-layout'});
      activeModal.componentInstance.programId = this.program.id;
      activeModal.componentInstance.modalHeader = 'Large Modal';
  }

  /**
   * Display All Sizes for one person
   */
  displayPersonsAllSizes() {
    if (this.program.person.id == null || this.program.person.id <= 0) {
      this.helperService.showWarning('شخص مورد نظر جهت نمایش اطلاعات سایز تعیین نشده است.');
      return;
    }
    const activeModal = this.modalService.open(DisplayAllSizesForOnePersonComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.personId = this.program.person.id;
    activeModal.componentInstance.personName = this.program.personName;
    activeModal.componentInstance.modalHeader = 'Large Modal';

    activeModal.result.then((resutl: any) => {
    }).catch(e => {
      // this.helperService.showError('error on modal edit program : ' + e)
    });
  }
}
