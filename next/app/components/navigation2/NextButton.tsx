"use client";
import { useRouter } from "next/navigation";
import { ChevronDownIcon } from "../icons/ChevronDownIcon";
import styles from "./NextButton.module.css";

interface Props {
  href: string;
}

export function NextButton(props: Props) {
  const router = useRouter();

  function onClick() {
    router.replace(props.href);
  }

  return (
    <button className={styles.component} onClick={onClick}>
      <div className={styles.smartphone}>
        <ChevronDownIcon />
      </div>
      <div className={styles.desktop}>
        <div>NEXT</div>
        <ChevronDownIcon />
      </div>
    </button>
  );
}
