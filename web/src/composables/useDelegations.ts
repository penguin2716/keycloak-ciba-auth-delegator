import { ref } from "vue";
import { delegationApi, type Delegation } from "@/api/delegations";

export function useDelegations() {
  const delegation = ref<Delegation | null>(null);
  const delegations = ref<Array<Delegation>>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);

  const listDelegations = async () => {
    loading.value = true;
    error.value = null;
    try {
      delegations.value = await delegationApi.list();
    } catch (err) {
      error.value = "Failed to get delegations";
    } finally {
      loading.value = false;
    }
  };

  const getDelegationById = async (id: string) => {
    loading.value = true;
    try {
      delegation.value = await delegationApi.getById(id);
    } catch (err) {
      error.value = "Failed to get delegation";
    } finally {
      loading.value = false;
    }
  };

  const approveDelegationById = async (id: string) => {
    loading.value = true;
    try {
      await delegationApi.approveById(id);
    } catch (err) {
      error.value = "Failed to approve delegation";
    } finally {
      loading.value = false;
    }
  };

  const cancelDelegationById = async (id: string) => {
    loading.value = true;
    try {
      await delegationApi.cancelById(id);
    } catch (err) {
      error.value = "Failed to cancel delegation";
    } finally {
      loading.value = false;
    }
  };

  const unauthorizeDelegationById = async (id: string) => {
    loading.value = true;
    try {
      await delegationApi.unauthorizeById(id);
    } catch (err) {
      error.value = "Failed to unauthorize delegation";
    } finally {
      loading.value = false;
    }
  };

  const deleteDelegationById = async (id: string) => {
    loading.value = true;
    try {
      await delegationApi.deleteById(id);
    } catch (err) {
      error.value = "Failed to delete delegation";
    } finally {
      loading.value = false;
    }
  };

  return {
    delegation,
    delegations,
    loading,
    error,
    listDelegations,
    getDelegationById,
    approveDelegationById,
    cancelDelegationById,
    unauthorizeDelegationById,
    deleteDelegationById,
  };
}
