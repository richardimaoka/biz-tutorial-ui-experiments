"use client";
import { usePathname, useRouter } from "next/navigation";
import { ChevronDownIcon } from "../icons/ChevronDownIcon";
import styles from "./NextButton.module.css";
import Link from "next/link";

interface Props {
  nextStep: string;
}

export function NextButton(props: Props) {
  const pathname = usePathname();
  console.log("NextButton");

  return (
    <Link href={pathname + "?step=" + props.nextStep}>
      <button className={styles.component}>
        <div className={styles.smartphone}>
          <div>next</div>
          <ChevronDownIcon />
        </div>
        <div className={styles.desktop}>
          <div>NEXT</div>
          <ChevronDownIcon />
        </div>
      </button>
    </Link>
  );
}
