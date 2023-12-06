"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { ChevronUpIcon } from "../icons/ChevronUpIcon";
import styles from "./PrevButton.module.css";

interface Props {
  prevStep: string;
  disabled?: boolean;
}

export function PrevButton(props: Props) {
  const pathname = usePathname();

  const smartPhoneStyle = props.disabled
    ? `${styles.smartphone} ${styles.disabled}`
    : styles.smartphone;

  const desktopStyle = props.disabled
    ? `${styles.desktop} ${styles.disabled}`
    : styles.desktop;

  return (
    <Link href={pathname + "?step=" + props.prevStep}>
      <button className={styles.component}>
        <div className={smartPhoneStyle}>
          <div>prev</div>
          <ChevronUpIcon />
        </div>
        <div className={desktopStyle}>
          <div>PREV</div>
          <ChevronUpIcon />
        </div>
      </button>
    </Link>
  );
}
