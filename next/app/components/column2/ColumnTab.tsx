"use client";

import styles from "./ColumnTab.module.css";
import { useRouter } from "next/navigation";
import { TabProperties } from "./tabTypes";
import { ColumnTabIcon } from "./ColumnTabIcon";

type Props = TabProperties;

export function ColumnTab(props: Props) {
  const router = useRouter();

  function onClick() {
    // need to use router.replace instaed of <Link> not to mess up the browser history
    router.replace(props.href);
  }

  const selectStyle = props.isSelected ? styles.selected : styles.unselected;
  const outerClassName = `${styles.component} ${selectStyle}`;

  return (
    <button className={outerClassName} onClick={onClick}>
      <span className={styles.smartphone}>
        <ColumnTabIcon name={props.name} />
      </span>
      <span className={styles.desktop}>{props.name}</span>
    </button>
  );
}
