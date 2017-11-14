#!/bin/bash
echo "POST http://localhost:8080/api/v1/tasks/" | vegeta attack -body todo.json -duration=30s -rate=1000 | tee results-POST.bin | vegeta report
