import styles from "./ProjectDir.module.css";

interface Props {
  projectDir?: string | null;
}

export function ProjectDir(props: Props) {
  return props.projectDir ? (
    <div className={styles.projectdir}>{props.projectDir}</div>
  ) : (
    <div />
  );
}
