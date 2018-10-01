import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {HelperService} from './helper.service';
import {ProgramView} from '../datamodel/ProgramView';

@Injectable({
  providedIn: 'root',
})

export class ProgramService {

  public PROGRAM_PICTURE_NAME: string = 'ProgramPicture';

  public GET_PROGRAM_LIST: string  = '/getPrograms';
  public GET_PROGRAM_EXERCISE_IMAGE: string  = '/getProgramExerciseImage';
  public GET_PROGRAM_EXERCISE_PDF: string  = '/getProgramExercisePDF';

  private GET_PERSON_PROGRAMS_ALL_SIZES: string  = '/getPersonProgramsAllSizes';
  private UPLOAD_PROGRAM_PICTURES: string  = '/uploadProgramPictures';
  public DOWNLOAD_PROGRAM_PICTURES_SHRINKED: string  = '/downloadProgramPicturesShrinked';
  public DOWNLOAD_PROGRAM_PICTURES_ORIGINAL: string  = '/downloadProgramPicturesOriginal';

  public SAVE_PROGRAM: string  = '/saveProgram';
  public DELETE_PROGRAM: string  = '/deleteProgram';

  constructor(private http: HttpClient,
              private helperService: HelperService) { }

  public getProgramList(personId: number) {
    return this.http.get<ProgramView[]>(this.helperService.SERVER_URL + this.GET_PROGRAM_LIST + '/' + personId);
  }

  public getProgramExerciseImage(id: number) {
    return this.http.get(
      this.helperService.SERVER_URL + this.GET_PROGRAM_EXERCISE_IMAGE + '/' + id + '');
  }

  public getProgramExercisePDF(id: number) {
    return this.http.get(
      this.helperService.SERVER_URL + this.GET_PROGRAM_EXERCISE_PDF + '/' + id + '');
  }

  public editProgram(program: ProgramView) {
    return this.http.post(this.helperService.SERVER_URL + this.SAVE_PROGRAM, program);
  }

  public deleteProgram(program: ProgramView) {
    return this.http.delete(this.helperService.SERVER_URL + this.DELETE_PROGRAM + '/' + program.id + '');
  }

  public addProgram(program: ProgramView) {
    return this.http.post(this.helperService.SERVER_URL + this.SAVE_PROGRAM, program);
  }

  public getPersonProgramsAllSizes(programId: number) {
    return this.http.get(this.helperService.SERVER_URL + this.GET_PERSON_PROGRAMS_ALL_SIZES + '/' +  programId);
  }

  uploadProgramPicture(personMobileNumber: string, programId: number, pictureName: string, formData: any) {
    return this.http.post(this.helperService.SERVER_URL +
        this.UPLOAD_PROGRAM_PICTURES + '/' + personMobileNumber + '/' + programId + '/' + pictureName, formData);
  }

}
