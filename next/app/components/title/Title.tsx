import styles from "./Title.module.css";

interface Props {
  phaseNum: number;
  title: string;
}

export function Title(props: Props) {
  return (
    <div className={styles.component}>
      <h1 className={styles.title}>
        フェーズ{props.phaseNum}
        <br />
        {props.title}
      </h1>
    </div>
  );
}
