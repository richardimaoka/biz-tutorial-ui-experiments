"use client";

import { TerminalIcon } from "../icons/TerminalIcon";
import styles from "./ColumnTab.module.css";
import { useRouter } from "next/navigation";

export interface Props {
  isSelected: boolean;
  name: string;
  href: string;
}

export function ColumnTab(props: Props) {
  const router = useRouter();

  function onClick() {
    router.replace(props.href);
  }

  const selectStyle = props.isSelected ? styles.selected : styles.unselected;
  const outerClassName = `${styles.component} ${selectStyle}`;

  return (
    <button className={outerClassName} onClick={onClick}>
      <span className={styles.smartphone}>
        <TerminalIcon />
      </span>
      <span className={styles.desktop}>{props.name}</span>
    </button>
  );
}
