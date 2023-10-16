import styles from "./ColumnContents.module.css";

interface Props {
  children: React.ReactNode;
}

export async function ColumnContents(props: Props) {
  return <div className={styles.component}>{props.children}</div>;
}
