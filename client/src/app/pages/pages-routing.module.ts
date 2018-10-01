import { RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';

import { PagesComponent } from './pages.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import {PersonsComponent} from './caa/components/persons/persons.component';
import {ExercisesComponent} from './caa/components/exercise/exercises.component';
import {MyloginComponent} from "../auth/mylogin.component";

const routes: Routes = [{
  path: '',
  component: PagesComponent,
  children: [
    {
      path: 'dashboard',
      component: DashboardComponent,
    },
    {
      path: 'persons',
      component: PersonsComponent,
    },
    {
      path: 'exercises',
      component: ExercisesComponent,
    },
    {
      path: 'login',
      component: MyloginComponent,
    },
    {
      path: '',
      redirectTo: 'dashboard',
      pathMatch: 'full',
    },
  ],
}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class PagesRoutingModule {
}
