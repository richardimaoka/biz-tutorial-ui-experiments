"use client";
import { usePathname, useRouter } from "next/navigation";
import { BackwardFastIcon } from "../icons/BackwardFast";
import styles from "./ButtonToInitialStep.module.css";
import Link from "next/link";

export function ButtonToInitialStep() {
  const pathname = usePathname();

  const smartPhoneStyle = styles.smartphone;
  const desktopStyle = styles.desktop;

  return (
    <Link href={pathname}>
      <button className={styles.component}>
        <div className={smartPhoneStyle}>
          <div>1st page</div>
          <BackwardFastIcon />
        </div>
        <div className={desktopStyle}>
          <div>1st page</div>
          <BackwardFastIcon />
        </div>
      </button>
    </Link>
  );
}
