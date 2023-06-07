export interface SelectOption<T = string> {
  value: T;
  label: string;
}

export enum SelectActionType {
  'OPEN',
  'CLOSE',
  'TOGGLE',
  'KEYDOWN',
  'BLUR',
  'CLICK',
  'DBLCLICK',
  'CHANGE',
  'SELECT',
  'MOUSEENTER',
  'RESET',
  'SET_VALUE',
  'SET_SELECTION',
}

export type SelectState<T = string> = {
  value: string;
  selection: string;
  isOpen: boolean;
  isEditing: boolean;
  currentIndex: number;
  items: SelectOption<T>[];
  defaultItems: SelectOption<T>[];
  lastKnownSelection: string;
};

export type SelectAction = {
  type: SelectActionType;
  payload?: unknown;
};