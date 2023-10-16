"use client";
import { PlayIcon } from "../icons/PlayIcon";
import styles from "./PlayButton.module.css";

interface Props {
  onClick: () => void;
}

export function PlayButton(props: Props) {
  return (
    <button className={styles.component} onClick={props.onClick}>
      <div className={styles.smartphone}>
        <PlayIcon />
      </div>
      <div className={styles.desktop}>
        <div>PLAY</div>
        <PlayIcon />
      </div>
    </button>
  );
}
