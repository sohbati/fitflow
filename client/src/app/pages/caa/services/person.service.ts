import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {HelperService} from './helper.service';
import {Person} from '../datamodel/Person';

@Injectable({
  providedIn: 'root',
})

export class PersonService {

  public GET_PERSON_LIST: string  = '/getPersons';
  public GET_PERSON_SHORT_LIST: string  = '/getPersonShortList';
  public SAVE_PERSON: string  = '/savePerson';
  public DELETE_PERSON: string  = '/deletePerson';

  constructor(private http: HttpClient,
              private helperService: HelperService) { }

  public getPersonList() {
    return this.http.get<Person[]>(this.helperService.SERVER_URL + this.GET_PERSON_LIST);
  }

  public getPersonShortList() {
    return this.http.get<Person[]>(this.helperService.SERVER_URL + this.GET_PERSON_SHORT_LIST);
  }

  public editPerson(person: Person) {
    return this.http.post(this.helperService.SERVER_URL + this.SAVE_PERSON, person);
  }

  public deletePerson(person: Person) {
    return this.http.delete(this.helperService.SERVER_URL + this.DELETE_PERSON + '/' + person.id + '');
  }

  public addPerson(person: Person) {
    return this.http.post(this.helperService.SERVER_URL + this.SAVE_PERSON, person);
  }
}
