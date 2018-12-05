import { Component } from '@angular/core';

import { MENU_ITEMS } from './pages-menu';
import {TranslateService} from '@ngx-translate/core';

@Component({
  selector: 'ngx-pages',
  template: `
    <ngx-sample-layout>
      <nb-menu [items]="menu"></nb-menu>
      <router-outlet></router-outlet>
    </ngx-sample-layout>
  `,
})
export class PagesComponent {
  menu = MENU_ITEMS;

  constructor(public translate: TranslateService) {
    translate.addLangs(['en', 'fa']);
    translate.setDefaultLang('fa');

    // const browserLang = translate.getBrowserLang();
    // translate.use(browserLang.match(/en|fa/) ? browserLang : 'fa');

  }


}
