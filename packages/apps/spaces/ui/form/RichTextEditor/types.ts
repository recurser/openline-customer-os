import { ReactExtensions, UseRemirrorReturn } from '@remirror/react';
import {
  BlockquoteExtension,
  BoldExtension,
  BulletListExtension,
  FontFamilyExtension,
  FontSizeExtension,
  HeadingExtension,
  ItalicExtension,
  NodeFormattingExtension,
  OrderedListExtension,
  StrikeExtension,
  UnderlineExtension,
    LinkExtension
} from 'remirror/extensions';
import { AnyExtension } from 'remirror';

export type BasicEditorExtentions =
  | ItalicExtension
  | BoldExtension
  | StrikeExtension
  | UnderlineExtension
  | OrderedListExtension
  | NodeFormattingExtension
  | BlockquoteExtension
  | BulletListExtension
  | FontFamilyExtension
  | FontSizeExtension
  | LinkExtension
  | HeadingExtension;

export type RemirrorProps<T extends AnyExtension> = UseRemirrorReturn<
  ReactExtensions<T>
>;