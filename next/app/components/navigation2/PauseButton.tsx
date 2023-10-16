"use client";
import { PauseIcon } from "../icons/PauseIcon";
import styles from "./PauseButton.module.css";

interface Props {
  onClick: () => void;
}

export function PauseButton(props: Props) {
  return (
    <button className={styles.component} onClick={props.onClick}>
      <div className={styles.smartphone}>
        <PauseIcon />
      </div>
      <div className={styles.desktop}>
        <div>PAUSE</div>
        <PauseIcon />
      </div>
    </button>
  );
}
