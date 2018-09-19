
import {Person} from './Person';
import {Exercise} from './Exercise';
import {PersonView} from "./PersonView";

export class ProgramView {
  id: number;
  programName: string;
  person: PersonView;
  programDate: Date;
  description: string = '';

  programExercise1Items: ProgramExerciseItem[];
  programExercise2Items: ProgramExerciseItem[];
  programExercise3Items: ProgramExerciseItem[];
  programExercise4Items: ProgramExerciseItem[];

  personAge: number;
  personTall: number;
  personWeight: number;
  personChest: number;
  personWaist: number;
  personAbdomen: number;
  personArm: number;
  personForeArm: number;
  personThigh: number;
  personShin: number;
  personButt: number;
  personFatPercentage: number;
  personFatWeight: number;
  personMuscleWeight: number;
  personScore: number;

  shamsiProgramDate: string;
  personName: string;
}

export class ProgramExerciseItem {
  idpk: number ;
  id: number;
  exerciseId: number;
  exerciseName: string;
  exerciseSet: number;
  exerciseRepeat: number;
  exerciseRepeatType: string;
  exerciseRepeatTypeDesc: string;
}
