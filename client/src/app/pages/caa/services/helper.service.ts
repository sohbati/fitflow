/**
 * Created by sohbati on 7/22/2018.
 */
import {Inject, Injectable} from "@angular/core";
import {NotificationsService, NotificationType} from "angular2-notifications";
import {FormBuilder, FormGroup} from "@angular/forms";


@Injectable({
  providedIn: 'root',
})
export class HelperService {
  constructor(@Inject(NotificationsService) private notificationService: NotificationsService,
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

  public SERVER_URL = 'http://localhost:9000';

  private showNotification(content: string, type: NotificationType) {
    const temp = this.form.getRawValue();
    const title = '';

    delete temp.title;
    delete temp.content;
    delete temp.type;
    this.notificationService.create(title, content, type, temp);
  }

/////////////////////////////////////////////////////////////////////////

  public showSuccess(body: string) {
    this.showNotification(body, NotificationType.Success);
  }

  public showError2(error: string, err: any) {
    if (err && err.error && err.error.message) {
      error = error + ' :  ' +  err.error.message;
    }
    this.showError(error);
  }

  public showError(err: any) {
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
    const charCodeZero: number = '۰'.charCodeAt(0);
    return numberAsStr.replace(/[۰-۹]/g, function (w: string) {
      return ( w.charCodeAt(0) - charCodeZero) + '';
    });
  }
}
