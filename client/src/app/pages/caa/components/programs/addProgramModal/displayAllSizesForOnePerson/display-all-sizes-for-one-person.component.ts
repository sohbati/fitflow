import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import {HelperService} from '../../../../services/helper.service';
import {ProgramService} from '../../../../services/program.service';

interface SearchData {
  item: string;
  value: any;
}


@Component({
  selector: 'ngx-display-all-sizes-for-one-person',
  templateUrl: './display-all-sizes-for-one-person.component.html',
  styleUrls: ['./display-all-sizes-for-one-person.component.css'],
})

export class DisplayAllSizesForOnePersonComponent implements OnInit {

  public personId: number;
  public personName: string = '';
  sizeList: string[][];
  constructor(private ngbActiveModal: NgbActiveModal ,
              private programService: ProgramService,
              private helperService: HelperService) {
  }
  ngOnInit() {
    this.programService.getPersonProgramsAllSizes(this.personId).subscribe( (result: any[]) => {
      this.sizeList = result;
    }, error => {
        this.helperService.showError(' error : ' + error);
      })
    }

  close() {
    this.ngbActiveModal.close();
  }

  }

