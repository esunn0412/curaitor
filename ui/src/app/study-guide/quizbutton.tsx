"use client";

import Link from "next/link";
import React from "react";
import { Flex, Progress, Tooltip } from "antd";

const QuizButton: React.FC = () => {
  return (
    <div className="space-y-4">
      <Link
        href="/quiz"
        className="flex-box border bg-gray-50 border-black font-medium text-sm h-50 px-10 flex items-left "
      >
        Take Quiz
      </Link>

      <Flex gap="small" vertical>
    <Flex gap="small" wrap>
      
      <Tooltip title="3 done / 3 in progress / 4 to do">
        <Progress percent={60} success={{ percent: 30 }} type="dashboard" />
      </Tooltip>
    </Flex>
  </Flex>
    </div>
  );
};

export default QuizButton;
