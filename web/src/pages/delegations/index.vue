<template>
  <v-container fluid class="mt-2 px-8">
    <v-row>
      <v-col>
        <div class="text-h5">Delegations</div>
      </v-col>
      <v-col cols="auto">
        <v-btn
          v-if="Object.values(checks).includes(true)"
          class="mr-3"
          color="error"
          prepend-icon="mdi-delete"
          @click="deleteSelected()"
          >delete</v-btn
        >
        <v-btn color="primary" prepend-icon="mdi-reload" @click="reload"
          >reload</v-btn
        >
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-table>
          <thead>
            <tr>
              <th style="width: 60px">
                <v-checkbox
                  v-model="bulkcheck"
                  color="primary"
                  density="compact"
                  hide-details
                  :indeterminate="indeterminate"
                  :disabled="items.length == 0"
                ></v-checkbox>
              </th>
              <th style="width: 80px">ID</th>
              <th style="width: 80px">Status</th>
              <th>ACR Values</th>
              <th>Binding Message</th>
              <th style="width: 120px">Consent</th>
              <th>Login Hint</th>
              <th style="width: 100px">Scope</th>
              <th style="width: 100px">Token</th>
              <th style="width: 120px">CreatedAt</th>
              <th style="width: 120px">UpdatedAt</th>
              <th style="width: 0px">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in items" :key="item.id">
              <th>
                <v-checkbox
                  v-model="checks[item.id]"
                  color="primary"
                  density="compact"
                  hide-details
                ></v-checkbox>
              </th>
              <td>{{ item.id.replace(/-.+/, "") }}</td>
              <td>
                <v-chip
                  v-if="item.status == 'PENDING'"
                  density="compact"
                  color="primary"
                  >{{ item.status }}</v-chip
                >
                <v-chip
                  v-if="item.status == 'SUCCEED'"
                  density="compact"
                  color="success"
                  >{{ item.status }}</v-chip
                >
                <v-chip
                  v-if="item.status == 'CANCELLED'"
                  density="compact"
                  color="warning"
                  >{{ item.status }}</v-chip
                >
                <v-chip
                  v-if="item.status == 'UNAUTHORIZED'"
                  density="compact"
                  color="error"
                  >{{ item.status }}</v-chip
                >
              </td>
              <td>{{ item.acr_values.length > 0 ? item.acr_values : "-" }}</td>
              <td>
                {{
                  item.binding_message.length > 0 ? item.binding_message : "-"
                }}
              </td>
              <td>{{ item.consent_required ? "Required" : "Not required" }}</td>
              <td>{{ item.login_hint }}</td>
              <td>{{ item.scope }}</td>
              <td>
                <v-dialog max-width="1200">
                  <template v-slot:activator="{ props: activatorProps }">
                    <v-btn v-bind="activatorProps" variant="text">view</v-btn>
                  </template>
                  <template v-slot:default="{ isActive }">
                    <v-card>
                      <v-card-title class="mt-3">
                        Auth token for {{ item.id }}
                      </v-card-title>
                      <v-card-text>
                        <v-container>
                          <v-row>
                            <v-col cols="6">
                              <div>Raw JWT</div>
                              <v-textarea
                                readonly
                                no-resize
                                hide-details
                                auto-grow
                                :modelValue="item.auth_token"
                              ></v-textarea>
                            </v-col>
                            <v-col cols="6">
                              <v-row>
                                <v-col>
                                  <div>JWT header</div>
                                  <v-textarea
                                    readonly
                                    no-resize
                                    hide-details
                                    auto-grow
                                    :modelValue="
                                      decode64(item.auth_token.split('.')[0])
                                    "
                                  ></v-textarea>
                                </v-col>
                              </v-row>
                              <v-row>
                                <v-col>
                                  <div>JWT payload</div>
                                  <v-textarea
                                    readonly
                                    no-resize
                                    hide-details
                                    auto-grow
                                    :modelValue="
                                      decode64(item.auth_token.split('.')[1])
                                    "
                                  ></v-textarea>
                                </v-col>
                              </v-row>
                            </v-col>
                          </v-row>
                        </v-container>
                      </v-card-text>
                    </v-card>
                  </template>
                </v-dialog>
              </td>
              <td>
                {{ item.created_at.replace(/\..+/, "").replace("T", "\n") }}
              </td>
              <td>
                {{ item.updated_at.replace(/\..+/, "").replace("T", "\n") }}
              </td>
              <td>
                <v-menu>
                  <template v-slot:activator="{ props }">
                    <v-badge
                      location="top right"
                      color="error"
                      dot
                      v-if="item.status == 'PENDING'"
                    >
                      <v-icon v-bind="props">mdi-dots-vertical</v-icon>
                    </v-badge>
                    <v-icon v-bind="props" v-if="item.status != 'PENDING'"
                      >mdi-dots-vertical</v-icon
                    >
                  </template>
                  <v-list>
                    <v-list-item
                      v-if="item.status == 'PENDING'"
                      prepend-icon="mdi-check"
                      title="Approve"
                      @click="approveById(item.id)"
                    ></v-list-item>
                    <v-list-item
                      v-if="item.status == 'PENDING'"
                      prepend-icon="mdi-cancel"
                      title="Cancel"
                      @click="cancelById(item.id)"
                    ></v-list-item>
                    <v-list-item
                      v-if="item.status == 'PENDING'"
                      prepend-icon="mdi-close"
                      title="Unauthorize"
                      @click="unauthorizeById(item.id)"
                    ></v-list-item>
                    <v-list-item
                      prepend-icon="mdi-delete"
                      title="Delete"
                      @click="deleteById(item.id)"
                    ></v-list-item>
                  </v-list>
                </v-menu>
              </td>
            </tr>
          </tbody>
        </v-table>
      </v-col>
    </v-row>

    <v-row v-if="items.length == 0">
      <v-col class="text-center">
        <div>No data found</div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { type Ref, ref, watch, onMounted } from "vue";

interface Delegation {
  id: string;
  status: string;
  acr_values: string;
  binding_message: string;
  consent_required: boolean;
  login_hint: string;
  scope: string;
  auth_token: string;
  created_at: string;
  updated_at: string;
}

const items: Ref<Delegation[]> = ref([]);

const bulkcheck: Ref<boolean> = ref(false);
const indeterminate: Ref<boolean> = ref(false);
const checks: Ref<Record<string, boolean>> = ref<Record<string, boolean>>({});

// 一括選択チェックボックスが変更されたときの挙動
watch(bulkcheck, async (newValue, oldValue) => {
  if (newValue) {
    checkAll();
    indeterminate.value = false;
  } else {
    uncheckAll();
    indeterminate.value = false;
  }
});

const checkAll = () => {
  for (let item of items.value) {
    checks.value[item.id] = true;
  }
};

const uncheckAll = () => {
  for (let item of items.value) {
    checks.value[item.id] = false;
  }
};

// 個別チェックボックスの状態に合わせて一括チェックボックスを調整する
watch(
  checks,
  async (newValue, oldValue) => {
    // true と false の両方があれば indeterminate 扱いとする
    indeterminate.value =
      Object.values(newValue).includes(true) &&
      Object.values(newValue).includes(false);

    // true しかなければ bulkcheck = true とする
    if (
      Object.values(newValue).includes(true) &&
      !Object.values(newValue).includes(false)
    ) {
      bulkcheck.value = true;
    }
    // false しかなければ bulkcheck = false とする
    if (
      !Object.values(newValue).includes(true) &&
      Object.values(newValue).includes(false)
    ) {
      bulkcheck.value = false;
    }
  },
  { deep: true },
);

const decode64 = (b64: string | undefined): string => {
  if (b64 === undefined) {
    return "undefined";
  }
  return JSON.stringify(JSON.parse(window.atob(b64)), null, 2);
};

const reload = async () => {
  fetch((import.meta.env.VITE_API_BASE_URL || "") + "/api/delegations")
    .then((resp) => resp.json())
    .then((resp) => {
      checks.value = {};
      for (let item of resp) {
        checks.value[item.id] = false;
      }
      bulkcheck.value = false;
      items.value = resp;
    });
};

const approveById = async (id: string) => {
  fetch(
    (import.meta.env.VITE_API_BASE_URL || "") +
      `/api/delegations/${id}/approve`,
    {
      method: "PUT",
    },
  ).then((resp) => {
    if (resp.ok) {
      reload();
    }
  });
};

const cancelById = async (id: string) => {
  fetch(
    (import.meta.env.VITE_API_BASE_URL || "") + `/api/delegations/${id}/cancel`,
    {
      method: "PUT",
    },
  ).then((resp) => {
    if (resp.ok) {
      reload();
    }
  });
};

const unauthorizeById = async (id: string) => {
  fetch(
    (import.meta.env.VITE_API_BASE_URL || "") +
      `/api/delegations/${id}/unauthorize`,
    {
      method: "PUT",
    },
  ).then((resp) => {
    if (resp.ok) {
      reload();
    }
  });
};

const deleteById = async (id: string) => {
  fetch((import.meta.env.VITE_API_BASE_URL || "") + `/api/delegations/${id}`, {
    method: "DELETE",
  }).then((resp) => {
    if (resp.ok) {
      reload();
    }
  });
};

const deleteSelected = async () => {
  for (let id of Object.keys(checks.value)) {
    if (checks.value[id] == true) {
      await fetch(
        (import.meta.env.VITE_API_BASE_URL || "") + `/api/delegations/${id}`,
        {
          method: "DELETE",
        },
      );
    }
  }
  reload();
};

onMounted(async () => {
  await reload();
});
</script>
