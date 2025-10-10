/**
 * Created by Reza on 15/08/2018.
 */
interface ExerciseItemsSubList {
  exerciseItemDesc: string;
  exerciseItemId: number;
  set: number;
  repeat: number;
  repeatType: string;
  repeatTypeDesc: string;

}

interface SelectedIExerciseItems {
  id: number;
  subListItems?: ExerciseItemsSubList[];
}
