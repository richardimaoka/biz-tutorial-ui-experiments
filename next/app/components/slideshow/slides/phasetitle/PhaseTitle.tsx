import styles from "./PhaseTitle.module.css";

interface Props {
  phaseNum: number;
  title: string;
}

export function PhaseTitle(props: Props) {
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
