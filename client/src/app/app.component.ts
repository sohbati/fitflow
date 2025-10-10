/**
 * @license
 * Copyright Akveo. All Rights Reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */
import { Component, OnInit } from '@angular/core';
import { AnalyticsService } from './@core/utils/analytics.service';
import {TranslateService} from '@ngx-translate/core';
import { globals } from '../environments/globals';

@Component({
  selector: 'ngx-app',
  template: `
    <simple-notifications></simple-notifications>
    <router-outlet></router-outlet>`,
})
export class AppComponent implements OnInit {

  constructor(private analytics: AnalyticsService,
              public translate: TranslateService) {
    translate.addLangs(globals.langs);
    translate.setDefaultLang('fa');

    const browserLang = translate.getBrowserLang();
    translate.use(browserLang.match(/en|fr/) ? browserLang : 'en');

  }

  ngOnInit() {
    this.analytics.trackPageViews();
  }
}
