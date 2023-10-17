"use client";
import { useRouter } from "next/navigation";
import styles from "./TerminalTab.module.css";

interface Props {
  name: string;
  href: string;
  isSelected: boolean;
}

export function TerminalTab(props: Props) {
  const router = useRouter();
  const selectStyle = props.isSelected ? styles.selected : styles.unselected;
  const componentClassName = `${styles.component} ${selectStyle}`;

  function onClick() {
    // need to use router.replace instaed of <Link> not to mess up the browser history
    router.replace(props.href);
  }

  return (
    <button className={componentClassName} onClick={onClick}>
      <span className={styles.desktop}>{props.name}</span>
    </button>
  );
}
