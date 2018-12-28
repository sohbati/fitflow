import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {HelperService} from './helper.service';
import {Person} from '../datamodel/Person';
import {ApiService} from '../../../myauth/service/api.service';

@Injectable({
  providedIn: 'root',
})

export class PersonService {

  public GET_PERSON_LIST: string  = '/getPersons';
  public GET_PERSON: string  = '/getPerson';
  public GET_FIND_BY_NAME_FAMILY_PHONE: string  = '/findByNameAndFamilyAndPhone';
  public GET_PERSON_SHORT_LIST: string  = '/getPersonShortList';
  public SAVE_PERSON: string  = '/savePerson';
  public DELETE_PERSON: string  = '/deletePerson';

  constructor(private http: HttpClient,
              private apiService: ApiService,
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

  public deletePerson(id: number) {
    return this.http.delete(this.helperService.SERVER_URL + this.DELETE_PERSON + '/' + id + '');
  }

  public getPersonById(id: number) {
    return this.http.get(this.helperService.SERVER_URL + this.GET_PERSON + '/' + id + '');
  }

  public findByNameFamilyPhone(searchStr: string) {
    return this.http.get(this.helperService.SERVER_URL + this.GET_FIND_BY_NAME_FAMILY_PHONE + '/' + searchStr + '');
  }

  /**
   * @Deprecated
   * @param person
   * @returns {Observable<Object>}
   */
  public addPerson(person: Person) {
    return this.http.post(this.helperService.SERVER_URL + this.SAVE_PERSON, person);
  }

  addOrUpdatePerson(formData: any, hasFile: boolean) {
    const save = hasFile ? 'WithImage' : 'WithOutImage';
    return this.http.post(this.helperService.SERVER_URL + this.SAVE_PERSON + save, formData);
  }
}
