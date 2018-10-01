import { NgModule } from '@angular/core';

import { TreeModule } from 'angular-tree-component';
import { PagesComponent } from './pages.component';
import { DashboardModule } from './dashboard/dashboard.module';
import { PagesRoutingModule } from './pages-routing.module';
import { ThemeModule } from '../@theme/theme.module';
import { MiscellaneousModule } from './miscellaneous/miscellaneous.module';
import {PersonsComponent} from './caa/components/persons/persons.component';
import {ExercisesComponent} from './caa/components/exercise/exercises.component';
import {HelperService} from './caa/services/helper.service';
import {PersonService} from './caa/services/person.service';
import {ProgramService} from './caa/services/program.service';
import {ProgramsComponent} from './caa/components/programs/programs.component';
import {AddProgramModalComponent} from './caa/components/programs/addProgramModal/add-program-modal.component';

import { Ng2SmartTableModule } from 'ng2-smart-table';
import { DisplayProgramExerciseImageComponent } from
      './caa/components/programs/addProgramModal/displayProgramExerciseImageComponent/display-program-exercise-image.component';
import {ShareButtonsModule} from 'ngx-sharebuttons';
import {HttpClientJsonpModule} from '@angular/common/http';
import {TemperatureDraggerComponent} from './dashboard/temperature/temperature-dragger/temperature-dragger.component';
import {TreeComponent} from './components/tree/tree.component';
import {AddExerciseItemModalComponent} from
  './caa/components/programs/addProgramModal/addExerciseItemComponent/add-exercise-item.component';
import {DxAutocompleteModule, DxTemplateModule} from 'devextreme-angular';
import {DisplayAllSizesForOnePersonComponent} from './caa/components/programs/addProgramModal/displayAllSizesForOnePerson/display-all-sizes-for-one-person.component';
import {AddPersonModalComponent} from './caa/components/persons/addpersonmodal/add-person-modal.component';
import {MyloginComponent} from '../auth/mylogin.component';
const PAGES_COMPONENTS = [
  PagesComponent,
];

@NgModule({
  imports: [
    PagesRoutingModule,
    ThemeModule,
    DashboardModule,
    MiscellaneousModule,
    Ng2SmartTableModule,

    HttpClientJsonpModule,
    ShareButtonsModule.forRoot(),
    TreeModule,

    DxAutocompleteModule,
    DxTemplateModule,
  ],
  providers: [ HelperService, PersonService, ProgramService],
  declarations: [
    ...PAGES_COMPONENTS,

    PersonsComponent,
    AddPersonModalComponent,
    ExercisesComponent,
    ProgramsComponent,
    AddProgramModalComponent,
    AddExerciseItemModalComponent,
    DisplayAllSizesForOnePersonComponent,
    DisplayProgramExerciseImageComponent,
    TemperatureDraggerComponent,
    TreeComponent,
    MyloginComponent,
  ],
  entryComponents: [
    AddProgramModalComponent,
    DisplayAllSizesForOnePersonComponent,
    AddPersonModalComponent,
    DisplayProgramExerciseImageComponent,
    ProgramsComponent,
    AddExerciseItemModalComponent,
  ],
})
export class PagesModule {
}
