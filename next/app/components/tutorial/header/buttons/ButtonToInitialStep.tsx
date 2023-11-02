"use client";
import { useRouter } from "next/navigation";
import { BackwardFastIcon } from "../../../icons/BackwardFast";
import styles from "./ButtonToInitialStep.module.css";

interface Props {
  href: string;
}

export function ButtonToInitialStep(props: Props) {
  const router = useRouter();

  function onClick() {
    // need to use router.replace instaed of <Link> not to mess up the browser history
    router.replace(props.href);
  }

  return (
    <button className={styles.component} onClick={onClick}>
      <BackwardFastIcon />
    </button>
  );
}
