"use client";
import { useRouter } from "next/navigation";
import { ChevronUpIcon } from "../icons/ChevronUpIcon";
import styles from "./PrevButton.module.css";

interface Props {
  href: string;
}

export function PrevButton(props: Props) {
  const router = useRouter();

  function onClick() {
    // need to use router.replace instaed of <Link> not to mess up the browser history
    router.replace(props.href);
  }

  return (
    <button className={styles.component} onClick={onClick}>
      <div className={styles.smartphone}>
        <div>prev</div>
        <ChevronUpIcon />
      </div>
      <div className={styles.desktop}>
        <ChevronUpIcon />
        <div>PREV</div>
      </div>
    </button>
  );
}
