"use client";

import { ReactNode } from "react";

interface Props {
  children: ReactNode;
}

export function Inner(props: Props) {
  return <>{props.children}</>;
}
