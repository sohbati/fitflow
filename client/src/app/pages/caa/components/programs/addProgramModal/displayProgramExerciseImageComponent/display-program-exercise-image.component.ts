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
        this.styles = 'width: 800px; height: 100%';
        this.imageBase64 =  this.helperService.BASE_64_IMAGE_PREFIX + result.content;
        this.shareImageURL = this.helperService.SERVER_URL +
          this.programService.GET_PROGRAM_EXERCISE_PDF + '/' + this.programId
      });
  }

  print() {
    const url = this.helperService.SERVER_URL + this.programService.GET_PROGRAM_EXERCISE_PDF + '/' + this.programId;
    window.open(url);

    // const a: HTMLLinkElement = document.createElement('A');
    // const filePath = url;
    // a.href = filePath;
    // // a.download = filePath.substr(filePath.lastIndexOf('/') + 1);
    // document.body.appendChild(a);
    // a.click();
    // document.body.removeChild(a);
  }

  close() {
    this.ngbActiveModal.close();
  }

  swipe() {
    // const largeImage = document.getElementById('largeImage');
    // largeImage.style.display = 'block';
    // largeImage.style.width = 200 + 'px';
    // largeImage.style.height = 200 + 'px';
    // const url = largeImage.getAttribute('src');
    // window.open(url, 'Image', 'width=largeImage.stylewidth,height=largeImage.style.height,resizable=1');
    window.open(this.helperService.SERVER_URL + '/shareProgramImage' + '/' + this.programId);
  }
}
