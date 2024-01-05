"use client";
import { usePathname, useRouter } from "next/navigation";
import { BackwardFastIcon } from "../icons/BackwardFast";
import styles from "./InitPageButton.module.css";
import Link from "next/link";

export function InitPageButton() {
  const pathname = usePathname();

  const smartPhoneStyle = styles.smartphone;
  const desktopStyle = styles.desktop;

  return (
    <Link href={pathname}>
      <button className={styles.component}>
        <div className={smartPhoneStyle}>
          <div className={styles.unselectable}>p. 1</div>
          <BackwardFastIcon />
        </div>
        <div className={desktopStyle}>
          <div className={styles.unselectable}>p. 1</div>
          <BackwardFastIcon />
        </div>
      </button>
    </Link>
  );
}
