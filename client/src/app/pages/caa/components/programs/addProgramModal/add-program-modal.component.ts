import { Component, OnInit } from '@angular/core';
import {HelperService} from '../../../services/helper.service';
import { LocalDataSource } from 'ng2-smart-table';
import {CompleterService, CompleterData, CompleterItem} from 'ng2-completer';
import {PersonService} from '../../../services/person.service';
import {Person} from '../../../datamodel/Person';
import {AddExerciseItemModalComponent} from './addExerciseItemComponent/add-exercise-item.component';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {ProgramService} from '../../../services/program.service';
import {ProgramExerciseItem, ProgramView} from '../../../datamodel/ProgramView';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import {ExerciseService} from '../../../services/exercise.service';
import {DisplayProgramExerciseImageComponent} from './displayProgramExerciseImageComponent/display-program-exercise-image.component';
import DataSource from 'devextreme/data/data_source';
import ArrayStore from 'devextreme/data/array_store';

/**
 * This is PersonList Entery Interface, used in dxAutoCompleter
 */
interface SearchData {
  item: string;
  value: any;
}

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

  /**
   * used in dxAutoCompleter
   * this code initialized again in init method  **/
  dataSource =  new DataSource({
      store: new ArrayStore({
        key: 'item',
        data: [],
      }),
    });

  /** ng2-smart-table source **/
  source: LocalDataSource = new LocalDataSource();
  /** TODO should remove **/
  public dataService: CompleterData;
  /** now it is not used but may in future i need it in dxAutoCompleter for [Person List]**/
  public searchData: SearchData[] = [];

  program: ProgramView;

  exerciseList: SelectedIExerciseItems[] = [];

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
        private exerciseService: ExerciseService,
        private personService: PersonService,
              private completerService: CompleterService) {
    this.dataService = completerService.local(this.searchData, 'value', 'item');
  }
  ngOnInit() {
    const personList = [];
    this.personService.getPersonShortList().subscribe((data: Person[]) => {
      data.forEach((p: Person, index, array) => {
        const pItem = p.firstName + ' ' + p.lastName + '(' + p.mobileNumber + ')';
        this.searchData.push({item : pItem, value : p.id});
        personList.push(pItem);
      });

      this.dataSource =  new DataSource({
        store: new ArrayStore({
          key: 'item',
          data: personList,
        }),
      });
    });

    if (this.program == null) {
      this.program = new ProgramView();
      this.program.programName = '';
      this.program.shamsiProgramDate = '';
      this.program.person = new Person();
      this.program.personName = '';
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
      this.program.personFatPercentage = 0;
      this.program.personFatWeight = 0;
      this.program.personMuscleWeight = 0;
      this.program.personScore = 0;

      this.exerciseList = [];
    }else {
      if (this.program.programExerciseItems == null || this.program.programExerciseItems.length <= 0) {
        this.exerciseList = [];
      }else {
        this.program.programExerciseItems.forEach((item: ProgramExerciseItem, index, array) => {
            this.addItemToList(item.id, item)
        });
        this.exerciseList.forEach((value, index, array) => {
          value.id  = index + 1;
        })
      }
      this.prepareGridRows(this.exerciseList);
    }
  }

  addItemToList(id: number, item: ProgramExerciseItem) {
    if (this.exerciseList == null || this.exerciseList.length === 0) {
      this.addNewItem(item);
    }else {
      let isnew: boolean = true;
      for (let i = 0; i < this.exerciseList.length; i++) {
        if (this.exerciseList[i].id === id) {
          isnew = false;
          const subItem = this.getSubListItem(item);
          this.exerciseList[i].subListItems.push(subItem);
        }
      }
      if (isnew) {
        this.addNewItem(item);
      }
    }
  }
  addNewItem(item: ProgramExerciseItem) {
    if (this.exerciseList == null) {
      this.exerciseList = [];
    }

    const subListItem: ExerciseItemsSubList = this.getSubListItem(item);
    const e: SelectedIExerciseItems = {
      id : item.id,
    }
    e.subListItems = [];
    e.subListItems.push(subListItem);
    this.exerciseList.push(e);
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
  addExerciseClick(event): void {
    const activeModal = this.modalService.open(AddExerciseItemModalComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.selectedExcercisesList = this.exerciseList;
    activeModal.componentInstance.modalHeader = 'Large Modal';

    activeModal.result.then((exerciseList: SelectedIExerciseItems[]) => {
      this.exerciseList = exerciseList;
      this.prepareGridRows(exerciseList);
    }).catch(e => {
      // this.helperService.showError('error on modal add program : ' + e)
    });
  }

  editExerciseClick(event) {
    const activeModal = this.modalService.open(AddExerciseItemModalComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.selectedExcercisesList = this.exerciseList;
    activeModal.componentInstance.modalHeader = 'Large Modal';

    activeModal.result.then((exerciseList: SelectedIExerciseItems[]) => {
      this.exerciseList = exerciseList;
      this.prepareGridRows(exerciseList);
    }).catch(e => {
     // this.helperService.showError('error on modal edit program : ' + e)
    });
  }

  prepareGridRows(exerciseList: SelectedIExerciseItems[]) {
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
    this.source.load(exerciseAsStringList);
  }

  personCompleterSelected(event: CompleterItem) {
    this.program.person.id = event.originalObject ?  event.originalObject.value : '' ;
  }

  prepareExerciseList() {
    if (this.exerciseList == null || this.exerciseList.length === 0) {
      return;
    }
    // clean container
    this.program.programExerciseItems = [];
    let id: number = 0;
    this.exerciseList.forEach((item, index, array) => {
      id++;
      item.subListItems.forEach((value, index2, array2) => {
        const ex: ProgramExerciseItem = new ProgramExerciseItem();
        ex.id = id;
        ex.exerciseId = value.exerciseItemId;
        ex.exerciseName = value.exerciseItemDesc;
        ex.exerciseSet = value.set;
        ex.exerciseRepeat = value.repeat;
        ex.exerciseRepeatType = value.repeatType;

        this.program.programExerciseItems.push(ex);
      })
    })
  }

  preparePersonId() {
    this.searchData.forEach((personSearchData: SearchData, index, array) => {
      if (this.program.personName === personSearchData.item) {
        this.program.person.id = personSearchData.value;
      }
    })
  }

  confirmClick() {
    this.prepareExerciseList();
    this.preparePersonId();

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
      const activeModal = this.modalService.open(DisplayProgramExerciseImageComponent, {windowClass : 'modalWindowClass', container: 'nb-layout'});
      activeModal.componentInstance.programId = this.program.id;
      activeModal.componentInstance.modalHeader = 'Large Modal';

  }
}
