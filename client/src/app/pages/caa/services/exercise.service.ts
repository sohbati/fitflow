import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {HelperService} from './helper.service';
import {Exercise} from '../datamodel/Exercise';

@Injectable({
  providedIn: 'root',
})

export class ExerciseService {

  public GET_PERSON_LIST: string  = '/getExercises';
  public GET_EXERCISE_SHORT_LIST: string  = '/getExerciseShortList';
  public SAVE_PERSON: string  = '/saveExercise';
  public DELETE_PERSON: string  = '/deleteExercise';

  public exerciseRepeatTypeDesc: string[] = ['ثانیه', 'دقیقه', 'تکرار'];

  constructor(private http: HttpClient,
              private helperService: HelperService) { }

  public getExerciseList() {
    return this.http.get<Exercise[]>(this.helperService.SERVER_URL + this.GET_PERSON_LIST);
  }

  public editExercise(exercise: Exercise) {
    return this.http.post(this.helperService.SERVER_URL + this.SAVE_PERSON, exercise);
  }

  public deleteExercise(exercise: Exercise) {
    return this.http.delete(this.helperService.SERVER_URL + this.DELETE_PERSON + '/' + exercise.id + '');
  }

  public addExercise(exercise: Exercise) {
    return this.http.post(this.helperService.SERVER_URL + this.SAVE_PERSON, exercise);
  }

  public getExerciseShortList() {
    return this.http.get<Exercise[]>(this.helperService.SERVER_URL + this.GET_EXERCISE_SHORT_LIST);
  }

  public getExerciseRepeatTypeDesc(repeatType: string) {
    return repeatType === 'S' ?  this.exerciseRepeatTypeDesc[0] :
      repeatType === 'M' ?  this.exerciseRepeatTypeDesc[1] :
        repeatType === 'C' ?  this.exerciseRepeatTypeDesc[2] : '';
  }
}
