import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import {ProgramService} from '../../../../services/program.service';
import {ImageView} from '../../../../datamodel/ImageView';
import {HelperService} from '../../../../services/helper.service';
import {CeiboShare} from 'ng2-social-share';

@Component({
  selector: 'ngx-display-program-exercise-image-modal',
  templateUrl: './display-program-exercise-image.component.html',
  styleUrls: ['./display-program-exercise-image.component.css'],
})

export class DisplayProgramExerciseImageComponent implements OnInit {


  programId: number;
  shareImageURL: string = '';
  imageBase64: string = '../../../../../../../assets/images/wait.gif';
  styles: string = 'width: 20px; height: 20px; object-fit: cover';
  constructor(private programService: ProgramService,
        private helperService: HelperService,
              private ngbActiveModal: NgbActiveModal ,
              ) {
  }
  ngOnInit() {
    this.styles = 'width: 20px; height: 20px';
      this.programService.getProgramExerciseImage(this.programId).subscribe((result: ImageView ) => {
        // this.styles = 'width: 595px; height: 842px';
        this.styles = 'width: 100%; height: 100%';
        this.imageBase64 =  'data:image/png;base64,' + result.content;
        this.shareImageURL = this.helperService.SERVER_URL + '/shareProgramImage/' + this.programId
      });
  }

  print() {
    let printContents, popupWin;
    // printContents = document.getElementById('print-section').innerHTML;
    popupWin = window.open('', '_blank', 'top=0,left=0,height=100%,width=100%');
    popupWin.document.open();
    popupWin.document.write(
      '<html><head><title></title><style>*{display: hidden;}img{display: block; width: 100%; height: 100%;}</style>' +
       '</head><body onload="window.print();window.close()">' +
      '<img src="' + this.shareImageURL + '"></img></body></html>');
    popupWin.document.close();
  }
}
