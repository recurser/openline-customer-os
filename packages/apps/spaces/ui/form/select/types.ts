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
  'CREATE',
  'DBLCLICK',
  'SET_EDITABLE',
  'CHANGE',
  'SELECT',
  'MOUSEENTER',
  'RESET',
  'SET_VALUE',
  'SET_SELECTION',
  'SET_INITIAL_ITEMS',
  'SET_DEFAULT_ITEMS',
  'SET_DEFAULT_SELECTION',
}

export type SelectState<T = string> = {
  value: string;
  selection: string;
  isOpen: boolean;
  isEditing: boolean;
  isCreating: boolean;
  canCreate: boolean;
  currentIndex: number;
  items: SelectOption<T>[];
  defaultItems: SelectOption<T>[];
  defaultSelection: string;
};

export type SelectAction = {
  type: SelectActionType;
  payload?: unknown;
};