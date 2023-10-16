import styles from "./Column.module.css";

interface Props {
  children: React.ReactNode;
}

export async function Column(props: Props) {
  return <div className={styles.component}>{props.children}</div>;
}
