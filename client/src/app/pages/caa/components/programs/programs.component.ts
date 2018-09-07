import {Component, OnInit} from '@angular/core';
import { LocalDataSource } from 'ng2-smart-table';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {AddProgramModalComponent} from './addProgramModal/add-program-modal.component';
import {ProgramService} from '../../services/program.service';
import {ProgramView} from '../../datamodel/ProgramView';
import {HelperService} from '../../services/helper.service';

@Component({
  selector: 'ngx-programs',
  styleUrls: ['./programs.component.scss'],
  templateUrl: './programs.component.html',
})
export class ProgramsComponent implements OnInit {

  settings = {
    mode: 'external',
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
      programName: {
        title: 'نام برنامه ',
        type: 'string',
        width: '20%',
      },
      personName: {
        title: 'نام فرد ',
        type: 'string',
        width: '20%',
      },
      shamsiProgramDate: {
        title: 'تاریخ برنامه',
        type: 'string',
        width: '15%',
      },
      description: {
        title: 'شرح برنامه',
        type: 'string',
        width: '45%',
      },
    },
  };
  private programList: ProgramView[] = [];
  source: LocalDataSource = new LocalDataSource();

  constructor(private modalService: NgbModal,
              private helperService: HelperService,
              private programService: ProgramService) { }

  ngOnInit() {
    this.initProgramList();
  }

  initProgramList() {
    this.programService.getProgramList().subscribe((data: ProgramView[]) => {
      this.programList = data;
      this.source.load(this.programList);
    });
  }

  addClick(event): void {
    const activeModal = this.modalService.open(AddProgramModalComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.modalHeader = 'Large Modal';
    activeModal.result.then((program: ProgramView) => {
        this.programList.push(program);
        this.source.load(this.programList);
    });
  }

  editClick(event): void {
    const activeModal = this.modalService.open(AddProgramModalComponent, { size: 'lg', container: 'nb-layout' });
    activeModal.componentInstance.program = event.data;
    activeModal.componentInstance.modalHeader = 'Large Modal';
    activeModal.result.then((program: ProgramView) => {
      for (let i = 0; i < this.programList.length; i++) {
        if (this.programList[i].id === program.id) {
          this.programList[i] = program;
          this.source.load(this.programList);
          break;
        }
      }
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
    this.programService.deleteProgram(event.data).subscribe((data: any) => {
      this.helperService.showSuccess('اطلاعات حذف گردید');
      let index = -1;
      for (let i = 0; i < this.programList.length; i++) {
        if (this.programList[i].id === event.data.id) {
          index = i;
          break;
        }
      }
      if (index >= 0) {
        this.programList.splice(index, 1);
        this.source.refresh();
      }
    }, (error) => {
      this.helperService.showError('در حذف برنامه ایرادی بوجود آمد لطفا با پشتیبانی سیستم تماس بگیرید')
    });
  }
}
