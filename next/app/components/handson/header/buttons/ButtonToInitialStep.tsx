"use client";
import { usePathname, useRouter } from "next/navigation";
import { BackwardFastIcon } from "../../../icons/BackwardFast";
import styles from "./ButtonToInitialStep.module.css";
import Link from "next/link";

export function ButtonToInitialStep() {
  const pathname = usePathname();

  return (
    <Link href={pathname}>
      <button className={styles.component}>
        <BackwardFastIcon />
      </button>
    </Link>
  );
}
