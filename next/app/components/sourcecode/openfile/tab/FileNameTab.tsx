import styles from "./FileNameTab.module.css";

interface Props {
  fileName: string;
  isNewFile?: boolean | null;
}

export const FileNameTab = (props: Props): JSX.Element => {
  const componentStyle = props.isNewFile
    ? `${styles.component} ${styles.new}`
    : `${styles.component} ${styles.default}`;

  return (
    <div className={componentStyle}>
      {props.fileName}
      {props.isNewFile && " (new file!)"}
    </div>
  );
};
