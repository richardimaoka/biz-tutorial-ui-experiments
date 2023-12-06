"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { ChevronDownIcon } from "../icons/ChevronDownIcon";
import styles from "./NextButton.module.css";

interface Props {
  nextStep: string;
  disabled?: boolean;
}

export function NextButton(props: Props) {
  const pathname = usePathname();

  const smartPhoneStyle = props.disabled
    ? `${styles.smartphone} ${styles.disabled}`
    : styles.smartphone;

  const desktopStyle = props.disabled
    ? `${styles.desktop} ${styles.disabled}`
    : styles.desktop;

  return (
    <Link href={pathname + "?step=" + props.nextStep}>
      <button className={styles.component}>
        <div className={smartPhoneStyle}>
          <div>next</div>
          <ChevronDownIcon />
        </div>
        <div className={desktopStyle}>
          <div>NEXT</div>
          <ChevronDownIcon />
        </div>
      </button>
    </Link>
  );
}
