import { NgModule } from '@angular/core';
import { TreeModule } from 'angular-tree-component';

import { PagesComponent } from './pages.component';
import { DashboardModule } from './dashboard/dashboard.module';
import { PagesRoutingModule } from './pages-routing.module';
import { ThemeModule } from '../@theme/theme.module';
import { MiscellaneousModule } from './miscellaneous/miscellaneous.module';
import {PersonsComponent} from './caa/components/persons/persons.component';
import {ExercisesComponent} from './caa/components/exercise/exercises.component';
import {SettingsComponent} from './caa/components/settings/settings.component';
import {HelperService} from './caa/services/helper.service';
import {PersonService} from './caa/services/person.service';
import {ProgramService} from './caa/services/program.service';
import {ProgramsComponent} from './caa/components/programs/programs.component';
import {AddProgramModalComponent} from './caa/components/programs/addProgramModal/add-program-modal.component';

import { Ng2SmartTableModule } from 'ng2-smart-table';
import { DisplayProgramExerciseImageComponent } from './caa/components/programs/addProgramModal/displayProgramExerciseImageComponent/display-program-exercise-image.component';
import {HttpClientJsonpModule} from '@angular/common/http';
import {TemperatureDraggerComponent} from './dashboard/temperature/temperature-dragger/temperature-dragger.component';
import {TreeComponent} from './components/tree/tree.component';
import {AddExerciseItemModalComponent} from './caa/components/programs/addProgramModal/addExerciseItemComp/add-exercise-item.component';
import {DxAutocompleteModule, DxTemplateModule} from 'devextreme-angular';
import {DisplayAllSizesForOnePersonComponent} from './caa/components/programs/addProgramModal/displayAllSizesForOnePerson/display-all-sizes-for-one-person.component';
import {AddPersonModalComponent} from './caa/components/persons/addpersonmodal/add-person-modal.component';
import {AddExerciseModalComponent} from './caa/components/programs/addProgramModal/addExerciseItemComp/addExerciseModalComponent/add-exercise-modal.component';

import {TranslateModule, TranslateLoader} from '@ngx-translate/core';
import {HttpClient} from '@angular/common/http';
import {TranslateHttpLoader} from '@ngx-translate/http-loader';

const PAGES_COMPONENTS = [
  PagesComponent,
];

// AoT requires an exported function for factories
export function HttpLoaderFactory(httpClient: HttpClient) {
  return new TranslateHttpLoader(httpClient);
}

@NgModule({
  imports: [
    PagesRoutingModule,
    ThemeModule,
    DashboardModule,
    MiscellaneousModule,
    Ng2SmartTableModule,

    HttpClientJsonpModule,
    TreeModule,

    DxAutocompleteModule,
    DxTemplateModule,
    TranslateModule.forRoot({
      loader: {
        provide: TranslateLoader,
        useFactory: HttpLoaderFactory,
        deps: [HttpClient],
      },
    }),
  ],
  providers: [ HelperService, PersonService, ProgramService],
  declarations: [
    ...PAGES_COMPONENTS,

    PersonsComponent,
    AddPersonModalComponent,
    ExercisesComponent,
    SettingsComponent,
    AddExerciseModalComponent,
    ProgramsComponent,
    AddProgramModalComponent,
    AddExerciseItemModalComponent,
    DisplayAllSizesForOnePersonComponent,
    DisplayProgramExerciseImageComponent,
    TemperatureDraggerComponent,
    TreeComponent,
  ],
  entryComponents: [
    AddProgramModalComponent,
    DisplayAllSizesForOnePersonComponent,
    AddPersonModalComponent,
    DisplayProgramExerciseImageComponent,
    ProgramsComponent,
    AddExerciseItemModalComponent,
    AddExerciseModalComponent,
  ],
})
export class PagesModule {
}
