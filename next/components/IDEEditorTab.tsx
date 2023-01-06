interface IDEEditorTabProps {
  filename: string;
}

export const IDEEditorTab = ({ filename }: IDEEditorTabProps): JSX.Element => {
  return <div>{filename}</div>;
};
