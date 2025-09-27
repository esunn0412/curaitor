"use client";

import { usePathname } from "next/navigation";
import { useEffect, useRef } from "react";
import * as vis from "vis-network";

const FileGraph = () => {
  const graphRef = useRef(null);
  const pathname = usePathname();
  const course = pathname.split("/")[1];

  useEffect(() => {
    if (!graphRef.current) return;
    const network = new vis.Network(
      graphRef.current,
      { nodes, edges },
      options,
    );
    network.on("beforeDrawing", (ctx: CanvasRenderingContext2D) => {
      if (course === "") return;
      nodes.forEach((node) => {
        const box = network.getBoundingBox(node.id!);
        if (node.label!.includes(course)) {
          ctx.fillStyle = "#fff085";
          ctx.roundRect(
            box.left,
            box.top,
            box.right - box.left,
            box.bottom - box.top,
            10,
          );
          ctx.fill();
        }
      });
    });
    network.redraw();
  }, [course]);

  return <div ref={graphRef} className="border basis-1/3"></div>;
};

export default FileGraph;

const nodes: vis.Node[] = [
  { id: 1, label: "cs-326/lectures/lecture-1.pdf" },
  { id: 2, label: "cs-326/notes/note-3.pdf" },
  { id: 3, label: "cs-253/notes/note-3.pdf" },
  { id: 4, label: "cs-170/lectures/lecture-11.pdf" },
  { id: 5, label: "cs-255/lectures/lecture-5.pdf" },
  { id: 6, label: "cs-224/lectures/lecture-5.pdf" },
];

const edges: vis.Edge[] = [
  { from: 1, to: 3 },
  { from: 1, to: 2 },
  { from: 2, to: 4 },
  { from: 2, to: 5 },
  { from: 3, to: 3 },
];

const options: vis.Options = {
  nodes: {
    value: 10,
    borderWidth: 1,
    color: {
      border: "#dadada",
      background: "#ffffff",
    },
    shape: "box",
    font: {
      face: "Geist Mono",
    },
    scaling: {
      min: 10,
      max: 30,
    },
  },
  interaction: {
    hover: true,
  },
};
