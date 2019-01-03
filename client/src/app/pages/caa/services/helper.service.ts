/**
 * Created by sohbati on 7/22/2018.
 */
import {Inject, Injectable} from '@angular/core';
import {NotificationsService, NotificationType} from 'angular2-notifications';
import {FormBuilder, FormGroup} from '@angular/forms';
import {Router} from '@angular/router';


@Injectable({
  providedIn: 'root',
})
export class HelperService {
  constructor(@Inject(NotificationsService) private notificationService: NotificationsService,
              @Inject(Router) private router: Router,
              @Inject(FormBuilder) private _fb: FormBuilder) {

    this.form = this._fb.group({
      type: 'success',
      title: 'This is just a title',
      content: 'This is just some content',
      timeOut: 5000,
      showProgressBar: true,
      pauseOnHover: true,
      clickToClose: true,
      animate: 'fromRight',
    });
  }

  form: FormGroup;

  // public SERVER_URL = 'http://localhost:9000';
  public SERVER_URL = '/api';
  public BASE_64_IMAGE_PREFIX: string = 'data:image/png;base64,';
  private NOT_LOGGED_IN_EXCEPTION = 'UserNotLoggedInException';
  private LOGIN_URL = '/myauth/mylogin';

  private showNotification(content: string, type: NotificationType) {
    const temp = this.form.getRawValue();
    const title = '';

    delete temp.title;
    delete temp.content;
    delete temp.type;
    this.notificationService.create(title, content, type, temp);
  }

/////////////////////////////////////////////////////////////////////////

  public showSuccess(body: string): void {
    this.showNotification(body, NotificationType.Success);
  }

  private isLogonError(err: any): boolean {
    if (err && err.error && err.error.exception && err.error.exception.indexOf(this.NOT_LOGGED_IN_EXCEPTION) >= 0) {
      this.router.navigate([this.LOGIN_URL]);
      return true;
    }
    return false;
  }

  public showError2(error: string, errObj: any) {
    if (this.isLogonError(errObj)) {
      return;
    }
    if (errObj && errObj.error && errObj.error.message) {
      error = error + ' :  ' +  errObj.error.message;
    }
    if (errObj && errObj.error && errObj.error.error && errObj.error.error.message) {
      error = error + ' :  ' +  errObj.error.error.message;
    }
    this.showError(error);
  }

  public showError(err: any) {
    if (this.isLogonError(err)) {
      return;
    }
    if (err && err.error && err.error.message) {
      this.showNotification(err.error.message, NotificationType.Error)
      return;
    }
    this.showNotification(err, NotificationType.Error);
  }

  public showWarning(body: string) {
    this.showNotification(body, NotificationType.Warn);
  }

  public convertToLatinNumbers(numberAsStr: string): string {
    if (numberAsStr == null || numberAsStr === '') {
      return '';
    }

    const charCodeZero: number = '۰'.charCodeAt(0);
    return numberAsStr.replace(/[۰-۹]/g, function (w: string) {
      return ( w.charCodeAt(0) - charCodeZero) + '';
    });
  }

  public  toInt(s: string): number {
    return parseInt(s, 10);
  }

  public toFloat(s: string): number {
    return parseFloat(s);
  }
}
