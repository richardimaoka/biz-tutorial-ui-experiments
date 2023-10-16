"use client";

import styles from "./ColumnTab.module.css";
import { useRouter } from "next/navigation";
import { TabName } from "./tabTypes";
import { ColumnTabIcon } from "./ColumnTabIcon";

function displayName(name: TabName): string {
  switch (name) {
    case "BackgroundImage":
      return "Background Image";
    case "Browser":
      return "Browser";
    case "DevTools":
      return "Dev Tools";
    case "ImageDescription":
      return "Image Description";
    case "Markdown":
      return "Markdown";
    case "SourceCode":
      return "Source Code";
    case "Terminal":
      return "Terminal";
    case "YouTube":
      return "YouTube";
  }
}

export type Props = {
  isSelected?: boolean;
  name: TabName;
  href: string;
};

export function ColumnTab(props: Props) {
  const router = useRouter();

  function onClick() {
    // need to use router.replace instaed of <Link> not to mess up the browser history
    router.replace(props.href);
  }

  const selectStyle = props.isSelected ? styles.selected : styles.unselected;
  const outerClassName = `${styles.component} ${selectStyle}`;
  const tabName = displayName(props.name);

  return (
    <button className={outerClassName} onClick={onClick}>
      <span className={styles.smartphone}>
        <ColumnTabIcon name={props.name} />
      </span>
      <span className={styles.desktop}>{tabName}</span>
    </button>
  );
}
