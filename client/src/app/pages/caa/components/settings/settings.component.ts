import { Component, OnInit } from '@angular/core';
import {AppSetting} from '../../datamodel/AppSetting';
import {AppSettingService} from '../../services/app-setting.service';
import {HelperService} from '../../services/helper.service';

@Component({
  selector: 'ngx-settings-list',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss'],
})

export class SettingsComponent implements OnInit {

  private appSetting: AppSetting = new AppSetting();

  constructor(
    private helperService: HelperService,
    private appSettingService: AppSettingService) {
  }

  ngOnInit() {
    this.appSettingService.getAppSetting().subscribe((data) => {
      this.mapData(data);
    }, err => {
      this.helperService.showError(err);
    });
  }

  mapData(data: any) {
    this.appSetting.config_isShowExerciseCodeInPrint = data.configItems.SHOW_CODE_IN_EXPORT_OR_PRINT === 'true';
  }

  save() {
    const data = new Object();
    data.configItems = new Object();
    data.configItems.SHOW_CODE_IN_EXPORT_OR_PRINT =
      this.appSetting.config_isShowExerciseCodeInPrint ? 'true' : 'false';

    this.appSettingService.save(data).subscribe( dd => {
      this.helperService.showSuccess('اطلاعات دخیره گردید');
    }, err => {
      this.helperService.showError(err);
    });

  }
}
