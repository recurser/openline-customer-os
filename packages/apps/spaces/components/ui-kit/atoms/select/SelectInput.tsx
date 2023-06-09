import React, { FC } from 'react';
import { useSelect } from '@spaces/atoms/select/useSelect';
import classNames from 'classnames';
import { InlineLoader } from '@spaces/atoms/inline-loader';
import styles from './select.module.scss';

export const SelectInput: FC<{ saving?: boolean; placeholder: string }> = ({
  saving,
  placeholder,
}) => {
  const { state, getInputProps, autofillValue } = useSelect();

  return (
    <>
      <span
        role='textbox'
        placeholder={placeholder}
        contentEditable={state.isEditing}
        className={classNames(styles.dropdownInput)}
        {...getInputProps()}
      />
      <span className={styles.autofill}>{autofillValue}</span>
      {saving && <InlineLoader />}
    </>
  );
};