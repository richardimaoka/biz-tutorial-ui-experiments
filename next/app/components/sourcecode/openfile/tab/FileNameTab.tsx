import styles from "./FileNameTab.module.css";

interface Props {
  fileName: string;
}

export const FileNameTab = (props: Props): JSX.Element => {
  return <div className={styles.component}>{props.fileName}</div>;
};
