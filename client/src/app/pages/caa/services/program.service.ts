import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {HelperService} from './helper.service';
import {ProgramView} from '../datamodel/ProgramView';

@Injectable({
  providedIn: 'root',
})

export class ProgramService {
  public GET_PROGRAM_LIST: string  = '/getPrograms';
  public GET_PROGRAM_EXERCISE_IMAGE: string  = '/getProgramExerciseImage';

  public SAVE_PROGRAM: string  = '/saveProgram';
  public DELETE_PROGRAM: string  = '/deleteProgram';

  constructor(private http: HttpClient,
              private helperService: HelperService) { }

  public getProgramList() {
    return this.http.get<ProgramView[]>(this.helperService.SERVER_URL + this.GET_PROGRAM_LIST);
  }

  public getProgramExerciseImage(id: number) {
    return this.http.get(
      this.helperService.SERVER_URL + this.GET_PROGRAM_EXERCISE_IMAGE + '/' + id + '');
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

}
